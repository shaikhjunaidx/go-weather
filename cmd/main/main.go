package main

import (
	"net/http"

	"github.com/shaikhjunaidx/go-weather/pkg/handlers"
)

func main() {
	http.HandleFunc("/welcome", handlers.WelcomeHandler)
	http.HandleFunc("/weather/", handlers.WeatherHandler)
	http.HandleFunc("/", handlers.FormHandler)

	http.ListenAndServe(":8080", nil)
}
