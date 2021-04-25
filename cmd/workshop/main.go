package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go_workshop_1/internal/config"
	"github.com/go_workshop_1/internal/handler"
	"gopkg.in/yaml.v2"
)

func main() {
	filename, err := filepath.Abs("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Server{}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("server shutdown")
}
