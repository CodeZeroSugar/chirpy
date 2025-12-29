package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/CodeZeroSugar/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("author_id")
	s := r.URL.Query().Get("sort")

	var chirps []database.Chirp
	var err error

	if len(id) == 0 {
		chirps, err = cfg.dbQueries.GetAllChirps(r.Context())
		if err != nil {
			log.Printf("Error fetching all chirps: %s", err)
			w.WriteHeader(404)
			return
		}
	} else {
		uuidString, err := uuid.Parse(id)
		if err != nil {
			log.Printf("Error parsing uuid while querying for all user's chirps: %v", err)
			w.WriteHeader(404)
		}
		chirps, err = cfg.dbQueries.GetAllChirpsForUserId(r.Context(), uuidString)
		if err != nil {
			log.Printf("Error getting chirps for userID from db, no records found: %v", err)
			respondWithError(w, 404, "chirps for provided author id were not found")
		}
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

	if s == "desc" {
		sort.Slice(fmtChirps, func(i, j int) bool {
			return fmtChirps[i].CreatedAt.After(fmtChirps[j].CreatedAt)
		})
	}
	respondWithJSON(w, 200, fmtChirps)
}
