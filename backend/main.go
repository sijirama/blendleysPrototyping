package main

import (
	"api/lib"
	_ "encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	// connect to the database
    lib.ConnectToDb()
    //defer db.Close()

	// router and server settings
	router := ServerSettings()

	log.Print("Server is starting...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
