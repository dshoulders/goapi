package services

import (
	"golang.org/x/crypto/bcrypt"
)

// Authenticate - Authenticate a user based on username and password
func Authenticate(username string, password string) (bool, error) {

	user, err := GetUser(username)

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
