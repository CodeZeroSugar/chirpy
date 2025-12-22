package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error with listen and server: %w", err)
	}
}
