package lib

import (
	"os"

	"github.com/gorilla/securecookie"
)

const (
	TOKEN_SECRET = "tokenfuckingsecret"
    JWT_COOKIE_NAME = "jwt_token"
)

func GetToken() string {
	var TOKEN = os.Getenv(os.Getenv("TOKEN_SECRET"))
	if len(TOKEN) == 0 {
		return TOKEN_SECRET
	}
	return TOKEN
}

func GetTokenSecret() *securecookie.SecureCookie {
	var hashKey = []byte("very-secret")
	var s = securecookie.New(hashKey, nil)
    return s
}
