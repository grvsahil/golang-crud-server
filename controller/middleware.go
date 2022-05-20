package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/grvsahil/projectEmployeeJS/logger"
	"github.com/grvsahil/projectEmployeeJS/model"
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

func CORS(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        origin := r.Header.Get("origin")
        w.Header().Set("Access-Control-Allow-Origin", origin)
        if r.Method == "OPTIONS" {
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
            
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Authorization")
            return
        } else {
            h.ServeHTTP(w, r)
        }
    }
}