package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		log.Printf("Error fetching all chirps: %s", err)
		w.WriteHeader(500)
		return
	}

	var fmtChirps []Chirp
	for _, chirp := range chirps {
		s := Chirp{
			ID:        chirp.UserID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
		fmtChirps = append(fmtChirps, s)
	}
	respondWithJSON(w, 200, fmtChirps)
}
