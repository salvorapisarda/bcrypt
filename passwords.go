package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	for i := 0; i < 100; i++ {
		password := "secret"
		hash, _ := HashPassword(password) // ignore error for the sake of simplicity

		fmt.Println("Password:", password)
		fmt.Println("Hash:    ", hash)

		match := CheckPasswordHash(password, hash)
		fmt.Println("Match:   ", match)

	}
	end := time.Now()
	elapsed := end.Sub(start).Seconds()
	fmt.Printf("time %s elapsed %d \n", time.Now(), elapsed)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
