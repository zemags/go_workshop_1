package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go_workshop_1/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("server shutdown")
}
