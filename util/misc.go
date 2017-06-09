package util

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(input string) string {
	data, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GetCurrentTimeMillis() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond))
}

func GetCurrentTimeSeconds() int64 {
	return time.Now().UnixNano() / (int64(time.Second))
}
