package controller

import (
	"net/http"
	"encoding/json"

	"golang-crud-server/logger"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	//get cookie from request
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Already logged out", http.StatusBadRequest)
		logger.ErrorLog.Println(err)
		return
	}

	//remove cookie
	cookie = &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	//send response
	err = json.NewEncoder(w).Encode("logged out")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}
