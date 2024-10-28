package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string, secret []byte) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func GenerateClaims(id string) jwt.MapClaims {
	return jwt.MapClaims{
		"Subject":   id,
		"ExpiredAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"Issuer":    "pelter",
	}
}

func GenerateToken(id uint, secret []byte) (string, error) {
	claims := GenerateClaims(fmt.Sprintf("%d", id))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
