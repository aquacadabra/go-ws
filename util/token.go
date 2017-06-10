package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

var ErrorTokenInvalid = errors.New("token: invalid")
var ErrorTokenExpired = errors.New("token: expired")

func GenerateToken(id int64, expiryInterval int64, secret string) string {

	// Create the Claims
	claims := TokenClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expiryInterval,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(secret))

	return ss
}

func ParseToken(tokenString string, secret string) (*TokenClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},

		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrorTokenInvalid
			}
			return []byte(secret), nil
		})

	if claims, ok := token.Claims.(*TokenClaims); ok &&
		err == nil &&
		token.Valid {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, ErrorTokenExpired
		}
		return claims, nil
	}

	return nil, ErrorTokenInvalid
}
