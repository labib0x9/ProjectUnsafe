package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(pass string, pepper string, cost int) (string, error) {
	pass += pepper
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(hash), err
}

func CompareHashAndPassword(passHash string, pass string, pepper string) bool {
	pass += pepper
	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(pass))
	return err == nil
}
