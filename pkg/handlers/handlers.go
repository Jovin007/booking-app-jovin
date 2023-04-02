package handlers

import (
	"booking-app-jovin/config"
	"booking-app-jovin/pkg/render"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	//fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("get-email.page.tmpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		//fmt.Println("Email-Id:", r.Form["email"])
		clientEmail := strings.Join(r.Form["email"], "")
		if IsEmailIdRegistered(clientEmail) {
			fmt.Println("Client Registered!!")

		} else {
			fmt.Println("Client Not Registered!!")
		}

	}
}

func IsEmailIdRegistered(email string) bool {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)
	rows, err := db.Query("SELECT COUNT(*) FROM client_details where email_id=$1", email)

	if err != nil {
		log.Fatal(err)

	}
	count := checkCount(rows)
	rows.Close()
	var isRegistered bool = false
	if count >= 1 {
		isRegistered = true
	}
	return isRegistered

}
func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
