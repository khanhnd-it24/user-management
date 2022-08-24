package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashAndPassword(hash []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	return hash, err
}
