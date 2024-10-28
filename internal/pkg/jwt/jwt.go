package jwt

import (
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

func GenerateToken(id uint, secret []byte) (string, error) {
	// create a new claims
	claims := jwt.MapClaims{
		"Subject":   id,
		"ExpiredAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"Issuer":    "pelter",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func GetIDFromToken(tokenString string, secret []byte) (string, error) {
	claims, err := ValidateToken(tokenString, secret)
	if err != nil {
		return "", err
	}
	return claims.(jwt.MapClaims)["Subject"].(string), nil
}
