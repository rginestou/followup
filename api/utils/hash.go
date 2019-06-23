package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash generates a salted hash for the input string
func GenerateHash(s string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// CompareWithHash compares string to generated hash
func CompareWithHash(hash string, s string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
}
