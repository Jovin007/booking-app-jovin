package handlers

import (
	"booking-app-jovin/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Authentication is the handler for the authentication page
func Authentication(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "authentication.page.tmpl")
}
