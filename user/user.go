package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	user string
	pass string
}

// hash a password from password bytes
func hash(pass []byte) (string, error) {
	h, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

// compare a hash with pass
func compare(hash string, pass []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, pass)
	if err != nil {
		return false
	}

	return true
}
