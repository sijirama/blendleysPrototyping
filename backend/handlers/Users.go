package handlers

import (
	"api/lib"
	_ "api/lib"
	Mytype "api/type"
	"encoding/json"
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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

func HandleLogInToAccount(w http.ResponseWriter, r *http.Request) {
	db := lib.DB
	var entry Mytype.UserLogin
	var u Mytype.Users
	json.NewDecoder(r.Body).Decode(&entry)

	err := db.QueryRow("SELECT * FROM users WHERE email = $1", entry.Email).
		Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Password, &u.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isMatch := lib.CheckPasswordHash(entry.Password, u.Password)
	if !isMatch {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Password dosent Match!")
		return
	}

	//json.NewEncoder(w).Encode("Matches")

	userClaim := lib.UserClaims{
		Id:    u.ID,
		First: u.FirstName,
		Last:  u.LastName,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token, err := lib.NewAccessToken(userClaim)
	if err != nil {
		log.Fatal("error creating access token")
	}

    s := lib.GetTokenSecret()
    encoded, err := s.Encode(lib.JWT_COOKIE_NAME, token)
    
	c := &http.Cookie{
		Name:  lib.JWT_COOKIE_NAME,
		Value: encoded,
        Path: "/",
        HttpOnly: true,
	}

    http.SetCookie(w, c)
	json.NewEncoder(w).Encode("Welcome to blendle!")
	return

}

func handleGetAccount(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
}
