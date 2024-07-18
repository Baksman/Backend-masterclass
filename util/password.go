package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword := []byte(password)
	res, err := bcrypt.GenerateFromPassword(hashedPassword, bcrypt.DefaultCost)
	return string(res), err
}

func CompareHashPassword(password string, hashedPassword string) error {
	bytePassowrd := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), bytePassowrd)
	return err
}
