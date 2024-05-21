package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/reveal.html"))

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
    
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/reveal", revealHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.ListenAndServe(":"+port, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func revealHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		answer := r.FormValue("answer")
		if answer == "Cain & Abel's" || answer == "cain & abel's" || answer == "Cain and Abel's" || answer == "cain and abel's" {
			templates.ExecuteTemplate(w, "reveal.html", nil)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
