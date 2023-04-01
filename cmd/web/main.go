package main

import (
	"booking-app-jovin/pkg/handlers"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const portNumber = ":8080"
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Premium@007"
	dbname   = "bookings"
)

// main is the main function
func main() {
	run()
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/get-email", handlers.GetEmail)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

func run() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!!")
}
