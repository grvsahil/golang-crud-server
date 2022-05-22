package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func Parse(tkn string) (Claims,error) {
	claims := Claims{}

	token, err := jwt.ParseWithClaims(tkn, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		return claims,err
	}

	if !token.Valid {
		err = errors.New("invalid token")
		return claims,err
	}

	return claims,nil
}
