package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/reveal.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/reveal", revealHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
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
