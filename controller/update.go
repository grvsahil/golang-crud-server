package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang-crud-server/db"
	"golang-crud-server/logger"
	"golang-crud-server/model"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var db = db.GetDatabase()

	//get id to update item
	vars := mux.Vars(r)
	id := vars["id"]
	Id, _ := strconv.Atoi(id)

	//check for already existing email
	var email string
	err := db.QueryRow("SELECT email FROM users where user_id=?", Id).Scan(&email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	//get new details from request
	var u model.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.ErrorLog.Println(err)
	}

	//perform validation
	valid := false
	if len(u.Firstname)+len(u.Lastname) < 30 && len(u.Email) < 20 {
		valid = true
	}
	if valid==false{
		http.Error(w,"Enter data of valid length",http.StatusBadRequest)
		logger.ErrorLog.Println("Email or password exceeds maximum length")
		return
	}

	//convert dob into required format
	dob, _ := time.Parse("2006-01-02", u.DOB)

	//check if new email is already in use 
	if email != u.Email && u.Email != "" {
		var countEmail int
		db.QueryRow("SELECT COUNT(email) FROM users where email=?", u.Email).Scan(&countEmail)
		if countEmail != 0 {
			http.Error(w, "Email already in use", http.StatusBadRequest)
			return
		}
	}

	//generate and execute query
	query := `update users SET updated_at="` + time.Now().Format("2006-01-02 15:04:05") + `"`
	if u.Firstname != "" {
		query += ` ,first_name = "` + u.Firstname + `"`
	}
	if u.Lastname != "" {
		query += ` ,last_name = "` + u.Lastname + `"`
	}
	if u.Email != "" {
		query += ` ,email = "` + u.Email + `"`
	}
	if u.DOB != "" {
		query += ` ,dob = "` + dob.Format("2006-01-02") + `"`
	}
	query += ` where user_id=` + id

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
		return
	}

	//send response
	err = json.NewEncoder(w).Encode("updated")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}
