package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"UserID"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Agent  string `json:"agent"`
	jwt.RegisteredClaims
}

var KEY = os.Getenv("JWT_KEY")

func GenerateToken(email, name, agent, id string) (string, error) {

	claims := Claims{
		Email:  email,
		Name:   name,
		UserID: id,
		Agent:  agent,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "user",
			ExpiresAt: jwt.NewNumericDate(jwt.NewNumericDate(time.Now()).AddDate(0, 0, 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(KEY))
}

func ValidateToken(token string) (*Claims, error) {

	claims := Claims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func RefreshToken(token string) (string, error) {

	claims := Claims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})

	if err != nil {
		return "", err
	}

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(jwt.NewNumericDate(time.Now()).AddDate(0, 0, 1))

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(KEY))

}
