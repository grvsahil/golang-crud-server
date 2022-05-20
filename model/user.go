package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email string
	jwt.StandardClaims
}

type Credentials struct {
	Email   string    `json:"email"`
	Password string `json:"password"`
}

type ReqUser struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Dob      string `json:"dob"`
	Password string `json:"password"`
}

type ResUser struct {
	Id    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Dob   string `json:"dob"`
}

type Paginate struct{
	Data []ResUser
	Total int
	Currpage int
}