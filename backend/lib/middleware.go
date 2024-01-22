package lib

import (
	"context"
	"net/http"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		//check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// pass to the next middleware
		next.ServeHTTP(w, r)
	})
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func JwtVerification(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := GetTokenSecret()
		if cookie, err := r.Cookie(JWT_COOKIE_NAME); err == nil {
			var value string
			if err = s.Decode(JWT_COOKIE_NAME, cookie.Value, &value); err == nil {

				user := ParseAccessToken(value)

				// Attach the user information to the request context
				ctx := context.WithValue(r.Context(), "user", user)
				//
				// // Create a new request with the updated context
				r = r.WithContext(ctx)
                

				h.ServeHTTP(w, r)
                return
			}
		}

		http.Error(w, "JWT not found or invalid", http.StatusUnauthorized)
	})
}
