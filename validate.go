package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
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
	respondWithJSON(w, 200, map[string]string{"cleaned_body": cleanedBody})
}
