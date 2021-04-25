package handler

import (
	"fmt"
	"net/http"

	"github.com/go_workshop_1/internal/api"
)

type Handler struct {
	weatherClient api.Client
}

// Dependency injection
func NewHandler(weatherClient api.Client) *Handler {
	return &Handler{
		weatherClient: weatherClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	weather, err := h.weatherClient.GetWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, weather.Weather)
}
