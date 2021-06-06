package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(unencryptedPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(unencryptedPassword), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func VerifyPassword(encryptedPassword string, unencryptedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(unencryptedPassword))

	if err != nil {
		return false
	}

	return true
}