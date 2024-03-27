package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	Email  string
	UserId int64
}

const jwtSecretKey = "secret" // TODO: use env variable

func GenerateJwtToken(payload JwtPayload) string {
	claims := jwt.MapClaims{
		"email":  payload.Email,
		"userId": payload.UserId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}
