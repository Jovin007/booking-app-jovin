package handlers

import (
	"booking-app-jovin/pkg/render"
	"fmt"
	"html/template"
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
func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("get-email.page.tmpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("Email-Id:", r.Form["email"])
	}
}
