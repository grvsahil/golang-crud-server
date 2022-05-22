package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"golang-crud-server/db"
	"golang-crud-server/logger"
	"golang-crud-server/model"
)

type Paginate struct {
	Data     []model.User
	Currpage int
	Total    int
	Lastpage int
}

func List(w http.ResponseWriter, r *http.Request) {
	var db = db.GetDatabase()

	//get query parameters from request
	archived := r.URL.Query().Get("archived")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	sortby := r.URL.Query().Get("sortby")
	order := r.URL.Query().Get("order")
	page := r.URL.Query().Get("page")
	items := r.URL.Query().Get("items")

	//make query string based on parameters for searching, sorting
	query := "select user_id,first_name,last_name,email,dob from users where archived=0"
	if archived == "true" {
		query = "select user_id,first_name,last_name,email,dob from users where archived=1"
	}
	if id != "" {
		query += " and user_id=" + id
	}
	if name != "" {
		query += ` and first_name like '%` + name + `%' or last_name like '%` + name + `%'`
	}
	if email != "" {
		query += ` and email like '%` + email + `%'`
	}
	if sortby != "" {
		if sortby=="name"{
			sortby = "first_name"
		}
		if sortby=="id"{
			sortby = "user_id"
		}
		if order != "" {
			query += ` order by ` + sortby + ` ` + order
		} else {
			query += ` order by ` + sortby + ` ASC`
		}
	}
	var total int
	user ,_:= db.Query(query)
	for user.Next(){
		total++
	}

	//perform pagination
	if items == "" {
		items = "3"
	}
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)
	i, _ := strconv.Atoi(items)
	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, i, (p-1)*i)

	//execute query
	user, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	//store retrieved data in structure
	var users []model.User
	for user.Next(){
		var u model.User
		user.Scan(&u.ID,&u.Firstname,&u.Lastname,&u.Email,&u.DOB)
		users =  append(users, u)
	}
	data := Paginate{
		Data:     users,
		Currpage: p,
		Total:    total,
		Lastpage: int(math.Ceil(float64(total)/float64(i))),
	}

	//send response
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}