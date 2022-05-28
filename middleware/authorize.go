package middleware

import (
	"net/http"

	"golang-crud-server/logger"
	"golang-crud-server/token"
)

func Authorize(hf http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set response header
		w.Header().Set("Content-Type", "application/json")

		//check relevant routes
		if (r.URL.Path == "/user"||r.URL.Path=="/users") && (r.Method == http.MethodGet ||
			r.Method==http.MethodPatch || r.Method == http.MethodDelete) {

			tkn := r.URL.Query().Get("token")

			//parse token string
			_, err := token.Parse(tkn)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				logger.ErrorLog.Println(err)
				return
			} else {
				hf.ServeHTTP(w, r)
				return
			}
		}
		
		hf.ServeHTTP(w, r)
	})
}
