package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		log.Fatalf("error when creating token : %v", err)
	}

	return tokenString
}
