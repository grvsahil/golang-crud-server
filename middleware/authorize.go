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
			
			//get cookie
			c, err := r.Cookie("token")
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

			//parse received cookie
			_, err = token.Parse(c.Value)
			if err != nil {
				logger.ErrorLog.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				hf.ServeHTTP(w, r)
			}
		}
		
		hf.ServeHTTP(w, r)
	})
}
