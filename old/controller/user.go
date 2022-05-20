package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang-crud-server/database"
	"golang-crud-server/logger"
	"golang-crud-server/model"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_pass")

func Login(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	var cred model.Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		logger.ErrorLog.Println(err)
	}

	var password string
	err = db.QueryRow("SELECT password FROM users where email=?", cred.Email).Scan(&password)
	if err != nil {
		http.Error(w, "Username or Password do not match", http.StatusUnauthorized)
		logger.ErrorLog.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(cred.Password))
	if err != nil {
		http.Error(w, "Username or Password do not match", http.StatusUnauthorized)
		logger.ErrorLog.Println(err)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 10)

	claims := &model.Claims{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	err = json.NewEncoder(w).Encode("login success")
	if err != nil {
		logger.ErrorLog.Println(err)
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Already logged out", http.StatusBadRequest)
		logger.ErrorLog.Println(err)
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
		logger.ErrorLog.Println(err)
	}
	if ValidateReq(&u) == false {
		http.Error(w, "Please check the entered data", http.StatusBadRequest)
		logger.ErrorLog.Println("Invalid data entered")
		return
	}

	dob, err := time.Parse("2006-01-02", u.Dob)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

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

	encPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	query := fmt.Sprintf(`INSERT INTO users (first_name,last_name,email,password,dob,created_at,archived) 
	VALUES ("%s", "%s", "%s","%s","%v","%v","%d")`, u.Fname, u.Lname, u.Email, string(encPass), dob.Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), 0)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
		return
	}

	json.NewEncoder(w).Encode("success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	vars := mux.Vars(r)
	id := vars["id"]
	Id, _ := strconv.Atoi(id)

	var email string
	err := db.QueryRow("SELECT email FROM users where user_id=?", Id).Scan(&email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	var u model.ResUser
	json.NewDecoder(r.Body).Decode(&u)
	if ValidateRes(&u) == false {
		http.Error(w, "Please check the entered data", http.StatusBadRequest)
		logger.ErrorLog.Println("Invalid data entered")
		return
	}
	dob, _ := time.Parse("2006-01-02", u.Dob)

	if Id != u.Id && u.Id != 0 {
		http.Error(w, "Cannot change Id once assigned", http.StatusBadRequest)
		return
	}
	if email != u.Email && u.Email != "" {
		var countEmail int
		db.QueryRow("SELECT COUNT(email) FROM users where email=?", u.Email).Scan(&countEmail)
		if countEmail != 0 {
			http.Error(w, "Email already in use", http.StatusBadRequest)
			return
		}
	}

	query := `update users SET updated_at="` + time.Now().Format("2006-01-02 15:04:05") + `"`
	if u.Fname != "" {
		query += ` ,first_name = "` + u.Fname + `"`
	}
	if u.Lname != "" {
		query += ` ,last_name = "` + u.Lname + `"`
	}
	if u.Email != "" {
		query += ` ,email = "` + u.Email + `"`
	}
	if u.Dob != "" {
		query += ` ,dob = "` + dob.Format("2006-01-02") + `"`
	}
	query += ` where user_id=` + id

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
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
	err := db.QueryRow("SELECT COUNT(user_id) FROM users where user_id=?", Id).Scan(&count)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}
	if count == 0 {
		http.Error(w, "Record not found", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE users set archived = 1 where user_id = %d;`, Id)
	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusForbidden)
		logger.ErrorLog.Println(err)
		return
	}
	json.NewEncoder(w).Encode("deleted")
}

func getQuery(r *http.Request) (string, int) {
	archived := r.URL.Query().Get("archived")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	sortby := r.URL.Query().Get("sortby")
	order := r.URL.Query().Get("order")
	page := r.URL.Query().Get("page")
	items := r.URL.Query().Get("items")

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
		if order != "" {
			query += ` order by ` + sortby + ` ` + order
		} else {
			query += ` order by ` + sortby + ` ASC`
		}
	}
	if items == "" {
		items = "3"
	}
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)
	i, _ := strconv.Atoi(items)
	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, i, (p-1)*i)
	return query, p
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var db = database.GetDatabase()
	query, page := getQuery(r)

	user, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		logger.ErrorLog.Println(err)
		return
	}

	var users []model.ResUser
	for user.Next() {
		var u model.ResUser
		user.Scan(&u.Id, &u.Fname, &u.Lname, &u.Email, &u.Dob)
		users = append(users, u)
	}

	var total int
	db.QueryRow("SELECT COUNT(user_id) FROM users where archived=0").Scan(&total)
	data := model.Paginate{
		Data:     users,
		Total:    total,
		Currpage: page,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
