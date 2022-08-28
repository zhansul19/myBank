package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("couldn't generate password: %w", err)
	}
	return string(hashPassword), err
}

func CheckPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
