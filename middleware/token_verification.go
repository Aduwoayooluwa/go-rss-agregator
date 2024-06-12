package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aduwoayooluwa/go-rss-scraper/utils"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// If no token is provided, return an error
		if tokenString == "" {

			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization token is required")
			//http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtKey, nil
		})

		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			// http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			// http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
