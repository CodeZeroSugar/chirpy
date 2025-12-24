package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/CodeZeroSugar/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	banned := []string{"kerfuffle", "sharbert", "fornax"}

	if len(params.Body) >= 140 {
		respondWithJSON(w, 400, map[string]string{"error": "Chirp too long"})
		return
	}
	split := strings.Split(params.Body, " ")
	for i, word := range split {
		if slices.Contains(banned, strings.ToLower(word)) {
			split[i] = "****"
		}
	}
	cleanedBody := strings.Join(split, " ")

	params.Body = cleanedBody

	chirpParams := database.CreateChirpParams{
		Body:   params.Body,
		UserID: params.UserID,
	}
	chirpDB, err := cfg.dbQueries.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, 500, "Error adding chirp to databse")
		return
	}
	c := Chirp{
		ID:        chirpDB.ID,
		CreatedAt: chirpDB.CreatedAt,
		UpdatedAt: chirpDB.UpdatedAt,
		Body:      chirpDB.Body,
		UserID:    chirpDB.UserID,
	}
	respondWithJSON(w, 201, c)
}
