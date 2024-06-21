package config

import (
	"encoding/json"
	"os"
)

type APIConfigData struct {
	OpenWeatherMapAPIKey string `json:"OpenWeatherMapApiKey"`
}

func LoadAPIKey(filename string) (APIConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return APIConfigData{}, err
	}

	var apiData APIConfigData
	err = json.Unmarshal(bytes, &apiData)
	if err != nil {
		return APIConfigData{}, err
	}

	return apiData, nil
}
