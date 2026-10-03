package core_encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPass), nil
}
func CompareHashAndPassword(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	); err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}
