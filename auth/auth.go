package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	// claims := &Claims{
	// Username: username,
	// StandardClaims: jwt.StandardClaims{
	// ExpiresAt: expirationTime.Unix(),
	// },
	// }

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func AuthenticateToken(tokenString string) error {
//
// }
//
