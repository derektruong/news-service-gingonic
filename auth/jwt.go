package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Create(id int, name, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = id
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}