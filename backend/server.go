package main

import (
	"api/handlers"
	"api/lib"
	"net/http"
	"github.com/gorilla/mux"
)

func ServerSettings() http.Handler {

	//create router
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/user", handlers.HandleCreateAccount).Methods("POST")

	//wrap
	enhancedRouter := lib.EnableCors(lib.JsonContentTypeMiddleware(router))

	return enhancedRouter
}
