package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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

	h := handler.NewHandler(apiClient, cfg.CustomWeather)

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}
	// handle shutdown gracefully
	// receive channel for incoming signal
	quit := make(chan os.Signal, 1)
	// read in the and to understand that program end down
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// goroutine wating for signal (os.Interrupt,
	// syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		// when received signal
		// when Background context is closing the child context close before
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		// ....
		done <- err
	}()

	log.Print("starting server")
	_ = srv.ListenAndServe()

	err = <-done

	log.Printf("server shutdown with %v", err)
}
