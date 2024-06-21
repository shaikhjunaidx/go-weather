package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/shaikhjunaidx/go-weather/pkg/config"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func Query(city string) (WeatherData, error) {
	apiConfig, err := config.LoadAPIKey(".apiConfig")
	if err != nil {
		return WeatherData{}, err
	}

	// URL-encode the city name
	encodedCity := url.QueryEscape(city)
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s", apiConfig.OpenWeatherMapAPIKey, encodedCity)

	response, err := http.Get(apiURL)
	if err != nil {
		return WeatherData{}, err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("API request failed with status %s", response.Status)
	}

	// Check the content type
	if contentType := response.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
		return WeatherData{}, fmt.Errorf("Invalid content type: %s", contentType)
	}

	var data WeatherData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return WeatherData{}, err
	}

	return data, nil
}
