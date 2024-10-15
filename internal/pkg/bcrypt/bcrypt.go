package bcrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(rawPassword string) (string, error) {
	rHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash a password")
	}
	return string(rHash), nil
}

func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
