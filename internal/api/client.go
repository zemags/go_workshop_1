package api

// Client interact with 3-rd party weather API
type Client interface {
	// GetWeather return weather
	GetWeather() (*WeatherResponse, error)
}
