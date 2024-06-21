package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
	"your-project/internal/weather"
)

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
		http.Redirect(w, r, "/weather/"+city, http.StatusSeeOther)
	}
}
