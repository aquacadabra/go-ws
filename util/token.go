package util

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

var ErrorTokenInvalid = errors.New("token: invalid")
var ErrorTokenExpired = errors.New("token: expired")

func ParseToken(tokenString string, secret string) (*TokenClaims, error) {
	return nil, ErrorTokenInvalid
}
