package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		log.Printf("Error getting api key from header: %v", err)
		respondWithError(w, 401, "upgrade request failed")
		return
	}
	if apiKey != cfg.polkaKey {
		log.Printf("Upgrade was attempted without correct api key")
		respondWithError(w, 401, "upgrade request failed")
		return
	}
	type dataParameters struct {
		UserID string `json:"user_id"`
	}
	type parameters struct {
		Event string         `json:"event"`
		Data  dataParameters `json:"data"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	if params.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}
	userUUID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		log.Printf("failed to parse userID while attempting to upgrade user: %v", err)
		respondWithError(w, 500, "something went wrong")
		return
	}
	_, err = cfg.dbQueries.UpgradeUser(r.Context(), userUUID)
	if err != nil {
		log.Printf("Attempted to upgrade user, could not find user: %v", err)
		respondWithError(w, 404, "user not found to upgrade")
		return
	}
	w.WriteHeader(204)
}
