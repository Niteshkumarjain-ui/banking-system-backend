package util

import (
	"banking-system-backend/domain"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (response domain.HashPassword, err error) {
	var hashPasswordBytes []byte
	hashPasswordBytes, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	response.EncryptPassword = string(hashPasswordBytes)
	return
}

func CheckPasswordHash(password, hash string) (response domain.CheckHashPassword) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	response.ValidPassword = err == nil
	return
}
