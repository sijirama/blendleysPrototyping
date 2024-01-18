package main

import (
	"api/handlers"
	"api/lib"
	"database/sql"
	_ "encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// connect to the databse
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	//create router
	router := mux.NewRouter()

	router.HandleFunc("/api/go/users", handlers.GetAllUsers(db)).Methods("GET")
	router.HandleFunc("/api/go/users", handlers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/api/go/users/{id}", handlers.GetUser(db)).Methods("GET")
	router.HandleFunc("/api/go/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/api/go/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

	//wrap
	enhancedRouter := lib.EnableCors(lib.JsonContentTypeMiddleware(router))

	//start server
	log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}
