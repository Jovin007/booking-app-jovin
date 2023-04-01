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

// GetEmail is the handler for the validation page
func GetEmail(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "get-email.page.tmpl")
}
