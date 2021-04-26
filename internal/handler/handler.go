package handler

import (
	"fmt"
	"net/http"

	"github.com/go_workshop_1/internal/api"
)

type Handler struct {
	weatherClient api.Client
	customWeather string
}

// Dependency injection
func NewHandler(weatherClient api.Client, customWeather string) *Handler {
	return &Handler{
		weatherClient: weatherClient,
		customWeather: customWeather,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	if h.customWeather != "" {
		fmt.Fprintf(w, h.customWeather)
		return
	}
	weather, err := h.weatherClient.GetWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, weather.Weather)
}
