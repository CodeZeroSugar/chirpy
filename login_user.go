package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLoginUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	user, err := cfg.dbQueries.FetchUser(r.Context(), params.Email)
	if err != nil {
		log.Printf("User lookup failed: %s", err)
		respondWithError(w, 401, "Incorrect email or password")
		return
	}
	result, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		log.Printf("Error comparing password to hash: %s", err)
		respondWithError(w, 401, "Incorrect email or password")
		return
	}
	if !result {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}
	userSuccess := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
	respondWithJSON(w, 200, userSuccess)
}
