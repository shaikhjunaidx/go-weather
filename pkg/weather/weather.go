package weather

import (
	"encoding/json"
	"net/http"
	"your-project/internal/config"
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

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?appid=" + apiConfig.OpenWeatherMapAPIKey + "&q=" + city)
	if err != nil {
		return WeatherData{}, err
	}

	defer response.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return WeatherData{}, err
	}

	return data, nil
}
