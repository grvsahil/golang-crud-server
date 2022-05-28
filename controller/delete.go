package controller

import (
	"net/http"
	"encoding/json"
	"strconv"

	"golang-crud-server/db"
	"golang-crud-server/logger"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var db = db.GetDatabase()

	//get id to delete item
	vars := mux.Vars(r)
	id := vars["id"]
	Id, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.ErrorLog.Println(err)
		return
	}

	//check for item in db
	var count int
	err = db.QueryRow("SELECT COUNT(user_id) FROM users where user_id=?", Id).Scan(&count)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}
	if count == 0 {
		http.Error(w, "Record not found", http.StatusBadRequest)
		return
	}

	//perform soft delete
	_, err = db.Exec("UPDATE users set archived = 1 where user_id=?",Id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
		return
	}

	//send response
	err = json.NewEncoder(w).Encode("deleted")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}
