package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Starting server on :8000")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
