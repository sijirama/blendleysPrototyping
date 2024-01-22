package handlers

import (
	"api/lib"
	_ "api/lib"
	Mytype "api/type"
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "log"
	"net/http"
)

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	db := lib.DB
	var u Mytype.Users
	json.NewDecoder(r.Body).Decode(&u)
	hashedPassword, _ := lib.HashPassword(u.Password)

	err := db.QueryRow("INSERT INTO users (firstname, lastname, username, password, email ) VALUES ($1 , $2, $3, $4, $5) RETURNING id", u.FirstName, u.LastName, u.Username, hashedPassword, u.Email).Scan(&u.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("User %s has been created with ID %v", u.FirstName, u.ID))
}

func handleLogInToAccount(w http.ResponseWriter, r *http.Request) {
}

func handleGetAccount(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
}
