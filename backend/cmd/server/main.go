package main

import (
	"log"
	"net/http"
	"url-shortener/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/shorten", handlers.ShortenHandler).Methods("POST")
	r.HandleFunc("/{shortURL}", handlers.RedirectHandler).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
