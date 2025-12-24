package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	pathParam := r.PathValue("chirpID")
	log.Println("chirpIDString:", pathParam)
	uuidPath, err := uuid.Parse(pathParam)
	if err != nil {
		log.Printf("Error parsing chirpID: %s", err)
		w.WriteHeader(500)
		return
	}
	chirp, err := cfg.dbQueries.GetChirp(r.Context(), uuidPath)
	if err != nil {
		respondWithError(w, 404, "Error: chirp ID does not exist")
		return
	}

	c := Chirp{
		ID:        chirp.UserID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
	respondWithJSON(w, 200, c)
}
