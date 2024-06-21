package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/shaikhjunaidx/go-weather/pkg/weather"
)

type WeatherPageData struct {
	City      string
	Celsius   float64
	Fahrenheit float64
	Error     string
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to GoLang weather project!"))
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := weather.Query(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/form.html")
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		city := r.FormValue("city")
		data, err := weather.Query(city)

		var pageData WeatherPageData
		if err != nil {
			pageData = WeatherPageData{Error: err.Error()}
		} else {
			kelvin := data.Main.Kelvin
			pageData = WeatherPageData{
				City:      data.Name,
				Celsius:   kelvin - 273.15,
				Fahrenheit: (kelvin-273.15)*9/5 + 32,
			}
		}

		tmpl, _ := template.ParseFiles("templates/form.html")
		tmpl.Execute(w, pageData)
	}
}
