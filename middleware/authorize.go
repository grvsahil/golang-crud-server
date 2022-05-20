package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang-crud-server/logger"
	"golang-crud-server/model"
)

func Authorize(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				logger.ErrorLog.Print(err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			logger.ErrorLog.Print(err)
			return
		}

		tokenStr := cookie.Value

		claims := &model.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				logger.ErrorLog.Print(err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			logger.ErrorLog.Print(err)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			logger.ErrorLog.Print(err)
			return
		}

		hf(w, r)
	}
}
