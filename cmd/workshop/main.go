package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go_workshop_1/internal/api/weather"
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

	apiClient := weather.NewWeatherClient(cfg.WeatherURL)

	h := handler.NewHandler(apiClient)

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)

	log.Print("starting server")
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("server shutdown")
}
