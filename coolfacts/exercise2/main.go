package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	factsRepo := repo{}
	factsRepo.add(fact{
		Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
		Description: "Did you know sonic is a hedgehog?!",
	})
	factsRepo.add(fact{
		Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
		Description: "You won't believe what happened to Arya!",
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "no http handler found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "text/plain")
		_, err := fmt.Fprint(w, "PONG")
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		// We are checking that we are answering only for GET  method
		if r.Method != http.MethodGet {
			http.Error(w, "no http handler found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "application/json")

		// Here we are getting the JSON encoding of for the all facts from the repo
		b, err := json.Marshal(factsRepo.getAll())
		if err != nil {
			errMessage := fmt.Sprintf("error marshaling facts : %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}

		// Next we are writing it the the http.ResponseWriter
		_, err = w.Write(b)
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})

	log.Println("starting server")
	err := http.ListenAndServe(":9003", nil)
	if err != nil {
		log.Fatal(err)
	}
}
