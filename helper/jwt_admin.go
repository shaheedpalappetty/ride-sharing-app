package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sub":      password,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
