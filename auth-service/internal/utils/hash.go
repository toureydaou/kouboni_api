package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("Password hash failed %v", err)
	}
	return string(hash)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Print(err)
	}
	return err == nil
}
