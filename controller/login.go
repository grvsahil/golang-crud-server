package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"golang-crud-server/db"
	"golang-crud-server/logger"
	"golang-crud-server/model"
	"golang-crud-server/token"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = db.GetDatabase()

	//get login credentials from request
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.ErrorLog.Println(err)
	}

	//match the credentials against records in database
	var password string
	err = db.QueryRow("SELECT user_id,password FROM users where email=?", user.Email).Scan(&password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.ErrorLog.Println(err)
		return
	}

	//generate password hash
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.ErrorLog.Println(err)
		return
	}

	//generate JWT token
	tokenString, err := token.GenToken(user.Email)
	if err != nil {
		logger.ErrorLog.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//set cookie for token
	expirationTime := time.Now().Add(time.Minute * 10)
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	//send response
	json.NewEncoder(w).Encode("login success")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}
