package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		log.Printf("Blocked attempt to reset without 'dev'")
		w.WriteHeader(403)
		return
	}
	cfg.fileserverHits.Store(0)
	msg := "Hits reset to 0\n"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))

	if err := cfg.dbQueries.ResetUsers(r.Context()); err != nil {
		log.Printf("Error creating user: %s", err)
		w.WriteHeader(500)
		return
	}
}
