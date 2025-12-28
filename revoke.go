package main

import (
	"log"
	"net/http"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("could not get token from header in revoke: %v", err)
		w.WriteHeader(500)
		return
	}
	refreshToken, err := cfg.dbQueries.FetchRefreshToken(r.Context(), bearerToken)
	if err != nil {
		log.Printf("no token in database to revoke: %v", err)
		w.WriteHeader(500)
		return
	}
	_, err = cfg.dbQueries.RevokeRefreshToken(r.Context(), refreshToken.Token)
	if err != nil {
		log.Printf("failed to revoke token: %v", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
