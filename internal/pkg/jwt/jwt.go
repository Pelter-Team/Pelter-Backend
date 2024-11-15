package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Token); ok {
		return claims, nil
	}
	return nil, errors.New("failed to validate token")
}

type Token struct {
	jwt.RegisteredClaims
	UserId uint
}

func GenerateToken(id uint) (string, error) {
	claims := Token{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "account",
			Issuer:    "pelter",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GetIDFromToken(tokenString string) (uint, error) {
	tokenClaims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, errors.New("get-id-from-token: cannot validate token")
	}

	fmt.Println(tokenClaims.UserId)

	return tokenClaims.UserId, nil
}
