package main

import (
	"net/http"
	"github.com/shaikhjunaidx/go-weather/pkg/handlers"
)

func main() {
	http.HandleFunc("/welcome", WelcomeHandler)
	http.HandleFunc("/weather/", WeatherHandler)
	http.HandleFunc("/", FormHandler)

	http.ListenAndServe(":8080", nil)
}
