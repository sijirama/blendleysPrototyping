package main

import (
	"api/handlers"
	"api/lib"
	"github.com/gorilla/mux"
	"net/http"
)

func ServerSettings() http.Handler {

	//create router
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/user/register", handlers.HandleCreateAccount).Methods("POST")
	router.HandleFunc("/api/user/authenticate", handlers.HandleLogInToAccount).Methods("POST")

    // secured routes
    securedRouter  := router.PathPrefix("").Subrouter()
    securedRouter.Use(lib.JwtVerification)
    securedRouter.HandleFunc("/api/home", handlers.WelcomeToBlendle).Methods("GET")

	//wrap
	enhancedRouter := lib.EnableCors(lib.JsonContentTypeMiddleware(router))

	return enhancedRouter
}
