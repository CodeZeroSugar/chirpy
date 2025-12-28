package main

import (
	"log"
	"net/http"

	"github.com/CodeZeroSugar/chirpy/internal/auth"
	"github.com/CodeZeroSugar/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("could not get token from header to delete chirp: %v", err)
		respondWithError(w, 401, "failed to delete chirp")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		log.Printf("could not get userID from jwt to delete chirp: %v", err)
		respondWithError(w, 401, "failed to delete chirp")
		return
	}

	pathParam := r.PathValue("chirpID")
	log.Println("chirpIDString:", pathParam)
	uuidChirp, err := uuid.Parse(pathParam)
	if err != nil {
		log.Printf("Error parsing chirpID: %s", err)
		w.WriteHeader(500)
		return
	}

	chirp, err := cfg.dbQueries.GetChirp(r.Context(), uuidChirp)
	if err != nil {
		log.Printf("chirp to be deleted was not in database: %v", err)
		respondWithError(w, 404, "chirp not found")
		return
	}

	if chirp.UserID != userID {
		log.Printf("chirps user ID did not match requesters user ID")
		w.WriteHeader(403)
		return
	}

	deleteParams := database.DeleteChirpParams{
		ID:     uuidChirp,
		UserID: userID,
	}
	if err = cfg.dbQueries.DeleteChirp(r.Context(), deleteParams); err != nil {
		log.Printf("failed to delete chirp from databse: %v", err)
		respondWithError(w, 500, "something went wrong")
		return
	}
	w.WriteHeader(204)
}
