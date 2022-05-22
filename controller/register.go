package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang-crud-server/db"
	"golang-crud-server/logger"
	"golang-crud-server/model"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var db = db.GetDatabase()

	//get user info from request
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.ErrorLog.Println(err)
		return
	}

	//perform validation
	var valid bool
	if len(u.Firstname)+len(u.Lastname) < 30 && (len(u.Password) > 8 && len(u.Password) < 20) && 
	len(u.Email) < 20 {
		valid = true
	}
	if valid==false{
		http.Error(w,"Enter data of valid length",http.StatusBadRequest)
		logger.ErrorLog.Println("Name, Email or password exceeds maximum length")
		return
	}

	//convert dob into required format
	dob, _ := time.Parse("2006-01-02", u.DOB)
	
	//match email to find any conflicts in records
	var countEmail int
	err = db.QueryRow("SELECT COUNT(email) FROM users where email=?", u.Email).Scan(&countEmail)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}
	if countEmail > 0 {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return
	}

	//generate password hash
	encPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	//execute query to store info into db
	query := fmt.Sprintf(`INSERT INTO users (first_name,last_name,email,password,dob,created_at,archived) 
	VALUES ("%s", "%s", "%s","%s","%v","%v","%d")`, u.Firstname, u.Lastname, u.Email, string(encPass),
	dob.Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), 0)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
		return
	}

	//send response
	err = json.NewEncoder(w).Encode("success")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}
