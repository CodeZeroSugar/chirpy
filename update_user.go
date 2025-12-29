package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
	"github.com/CodeZeroSugar/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("tried to update user with no access token: %v", err)
		respondWithError(w, 401, "Could not update user properties")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		log.Printf("failed to get user from token while updating user: %v", err)
		respondWithError(w, 401, "Could not update user properties")
		return
	}
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("failed to hash password when updating user properties: %s", err)
		respondWithError(w, 401, "Could not update user properties")
		return
	}
	updateUserParams := database.UpdateUserParams{
		ID:             userID,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	}
	user, err := cfg.dbQueries.UpdateUser(r.Context(), updateUserParams)
	if err != nil {
		log.Printf("failed to update user properties with new email and password: %v", err)
		respondWithError(w, 401, "Could not update user properties")
		return
	}
	updatedUser := User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}
	respondWithJSON(w, 200, updatedUser)
}
