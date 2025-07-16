package utils

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("my_super_secret_key_12345") // Use secure key in production

func GenerateJWT(userID uint) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
