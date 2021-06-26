package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Handle the root rout
func Home(w http.ResponseWriter, r *http.Request) {
	parsedHome, _ := template.ParseFiles("templates/home.page.tmpl", "templates/base.layout.tmpl")
	err := parsedHome.Execute(w, nil)
	if err != nil {
		log.Panicln("Error rendering home page: ", err)
	}
}

// Handle the /about route
func About(w http.ResponseWriter, r *http.Request) {
	parsedAbout, _ := template.ParseFiles("templates/about.page.tmpl", "templates/base.layout.tmpl")
	err := parsedAbout.Execute(w, nil)
	if err != nil {
		log.Panicln("Error rendering about page: ", err)
	}
}
