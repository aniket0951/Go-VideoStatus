package utils

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var latters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(size int) string {
	b := make([]rune, size)

	for i := range b {
		b[i] = latters[rand.Intn(len(latters))]
	}

	return string(b)
}

func generateRandomNumber() string {
	t := fmt.Sprint(time.Now().Nanosecond())
	tm := t[:7]
	return tm
}

func RandomMobileNumber() string {
	return generateRandomNumber() + "198"
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(14))
}



func HasAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}