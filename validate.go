package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	type returnValid struct {
		Valid bool `json:"valid"`
	}

	type returnError struct {
		Error string `json:"error"`
	}

	var respValid returnValid
	var respError returnError

	if len(params.Body) <= 140 {
		respValid.Valid = true
		dat, err := json.Marshal(respValid)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(500)
			w.Write([]byte(`{"error":"Something went wrong"`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(dat)

	} else {
		respError.Error = "Chirp is too long"
		dat, err := json.Marshal(respError)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(500)
			w.Write([]byte(`{"error":"Something went wrong"`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)
	}
}
