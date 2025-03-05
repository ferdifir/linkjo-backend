package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint, email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "linkjo"
	}

	claims := jwt.MapClaims{
		"tenant_id": userID,
		"email":     email,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
