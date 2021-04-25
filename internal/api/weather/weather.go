package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go_workshop_1/internal/api"
)

const getWeatherPath = "/api?format=json"

// WeatherClient is a weather API client
type WeatherClient struct {
	url string
}

func (wc *WeatherClient) GetWeather() (*api.WeatherResponse, error) {
	urlPath := wc.url + getWeatherPath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", http.StatusText(resp.StatusCode))
	}

	var data api.WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
