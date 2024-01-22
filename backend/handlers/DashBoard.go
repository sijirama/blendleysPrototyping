package handlers

import (
	"api/lib"
	"fmt"
	"net/http"
)

func WelcomeToBlendle(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*lib.UserClaims)
    fmt.Println("The value is: ",r.Context().Value("user"))
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	// Use the user information as needed
	fmt.Printf("Hello, %s!", user.Email)
}
