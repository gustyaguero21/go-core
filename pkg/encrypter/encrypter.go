package encrypter

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypter(password string) ([]byte, error) {
	createHash, createErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if createErr != nil {
		return nil, errors.New("error creating password hash")
	}
	return createHash, nil
}

func PasswordDecrypter(hash []byte, password string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return false
	}
	return true
}
