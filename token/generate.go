package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct{
	Email string
	jwt.StandardClaims
}

var key = []byte("my_pass")

func GenToken(email string) (string,error) {
	expirationTime := time.Now().Add(time.Minute * 10)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknString, err := token.SignedString(key)
	if err != nil {
		return tknString,err
	}

	return tknString,err
}
