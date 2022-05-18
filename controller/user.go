package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/grvsahil/projectEmployeeJS/database"
	"github.com/grvsahil/projectEmployeeJS/model"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_pass")

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	var cred model.Credentials
	json.NewDecoder(r.Body).Decode(&cred)
	var password string
	err := db.QueryRow("SELECT password FROM users where user_id=?", cred.UserId).Scan(&password)
	if err != nil {
		http.Error(w, "Username or Password do not match", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(cred.Password))
	if err != nil {
		http.Error(w, "Username or Password do not match", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &model.Claims{
		Id: cred.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	json.NewEncoder(w).Encode("login success")
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Already logged out", http.StatusBadRequest)
		return
	}

	cookie = &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("logged out")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	var u model.ReqUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}
	if Validate(&u) == false {
		http.Error(w,"Please check the entered data",http.StatusBadRequest)
		return
	}

	dob, err := time.Parse("2006-01-02", u.Dob)
	if err != nil {
		http.Error(w,"Internal server error",http.StatusInternalServerError)
		return
	}

	var count int
	err = db.QueryRow("SELECT COUNT(user_id) FROM users where user_id=?",u.Id).Scan(&count)
	if err != nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if count==1{
		http.Error(w, "UserId already taken", http.StatusBadRequest)
		return
	}

	encPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	query := fmt.Sprintf(`INSERT INTO users (user_id,first_name,last_name,email,password,dob,created_at,archived) 
	VALUES ("%d", "%s", "%s", "%s","%s","%v","%v","%d")`, u.Id, u.Fname, u.Lname, u.Email, string(encPass), dob.Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), 0)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode("success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	vars := mux.Vars(r)
	id := vars["id"]
	Id, _ := strconv.Atoi(id)

	var count int
	err := db.QueryRow("SELECT COUNT(user_id) FROM users where user_id=?",Id).Scan(&count)
	if err != nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if count==0{
		http.Error(w, "Record not found", http.StatusBadRequest)
		return
	}

	var u model.ResUser
	json.NewDecoder(r.Body).Decode(&u)
	dob, _ := time.Parse("2006-01-02", u.Dob)

	if Id!=u.Id{
		http.Error(w, "Cannot change Id once assigned", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE users
	SET first_name = "%s", last_name= "%s", email="%s", dob="%v", updated_at="%v"
	WHERE user_id = %d;`, u.Fname, u.Lname, u.Email, dob.Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), Id)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode("updated")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	vars := mux.Vars(r)
	id := vars["id"]
	Id, _ := strconv.Atoi(id)

	var count int
	err := db.QueryRow("SELECT COUNT(user_id) FROM users where user_id=?",Id).Scan(&count)
	if err != nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if count==0{
		http.Error(w, "Record not found", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE users set archived = 1 where user_id = %d;`, Id)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode("deleted")
}

func getQuery(r *http.Request) string {
	archived := r.URL.Query().Get("archived")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	sortby := r.URL.Query().Get("sortby")
	order := r.URL.Query().Get("order")

	query := "select user_id,first_name,last_name,email,dob from users where archived=0"
	if archived=="true"{
		query = "select user_id,first_name,last_name,email,dob from users where archived=1"
	}
	if id!=""{
		query += " and user_id="+id
	}
	if name!=""{
		query += ` and first_name like '%`+name+`%' or last_name like '%`+name+`%'`
	}
	if email!=""{
		query += ` and email like '%`+email+`%'`
	}
	if sortby!=""{
		if order!=""{
			query += ` order by `+sortby+` `+order
		}else{
			query += ` order by `+sortby+` ASC`
		}
	}
	return query
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	query := getQuery(r)

	user, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var users []model.ResUser
	for user.Next() {
		var u model.ResUser
		user.Scan(&u.Id, &u.Fname, &u.Lname, &u.Email, &u.Dob)
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
