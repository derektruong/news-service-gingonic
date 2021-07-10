package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/derektruong/news-app-gin/models"

	jwt "github.com/dgrijalva/jwt-go"
)

type CustomClaim struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func Create(id int, name, email string) (string, error) {

	claims := &CustomClaim{
		Uid:   strconv.Itoa(id),
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
			Issuer: "/",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func VerifyJWT(jwtToken string) (*models.Account, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&CustomClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_JWT")), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("jwt is expired")
	}

	id, err := strconv.Atoi(claims.Uid)

	if err != nil {
		return nil, err
	}

	return models.NewAccount(id, claims.Name, claims.Email, ""), nil
}
