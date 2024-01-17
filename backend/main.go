package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

    router.HandleFunc("/api/go/users" , getUsers(db) ).Methods("GET")
    router.HandleFunc("/api/go/users" , createUsers(db) ).Methods("POST")
    router.HandleFunc("/api/go/users/{id}" , getUsers(db) ).Methods("GET")
    router.HandleFunc("/api/go/users/{id}" , updateUsers(db) ).Methods("PUT")
    router.HandleFunc("/api/go/users/{id}" , deleteUsers(db) ).Methods("DELETE")
}
