package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang-crud-server/logger"
	"golang-crud-server/model"
)

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
