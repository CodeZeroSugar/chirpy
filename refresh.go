package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("could not get refresh token from header: %v", err)
		w.WriteHeader(500)
		return
	}
	refreshToken, err := cfg.dbQueries.FetchRefreshToken(r.Context(), token)
	if err != nil {
		log.Printf("could not fetch refresh token from database: %v", err)
		respondWithError(w, 401, "refresh token is expired or does not exists")
		return
	}
	user, err := cfg.dbQueries.GetUserFromRefreshToken(r.Context(), refreshToken.Token)
	if err != nil {
		log.Printf("could not get user from refresh token: %v", err)
		respondWithError(w, 401, "refresh token is expired or does not exists")
		return
	}
	expiresIn := time.Hour
	newToken, err := auth.MakeJWT(user.UserID, cfg.secret, expiresIn)
	if err != nil {
		log.Printf("failed to make new jwt during refresh: %v", err)
		w.WriteHeader(500)
		return
	}
	resp := User{
		Token: newToken,
	}

	respondWithJSON(w, 200, resp)
}
