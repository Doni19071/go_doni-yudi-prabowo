package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_JWT = "123"

func CreateToken(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}
