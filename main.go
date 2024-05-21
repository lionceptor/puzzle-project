package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("index.html", "reveal.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/reveal", revealHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func revealHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		answer := r.FormValue("answer")
		if answer == "The Tower" || answer == "the tower" || answer == "Tower" || answer == "tower" {
			templates.ExecuteTemplate(w, "reveal.html", nil)
			return
		}
		templates.ExecuteTemplate(w, "index.html", "Sorry, that’s not correct. Try again!")
	}
}
