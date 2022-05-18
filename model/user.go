package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Id int
	jwt.StandardClaims
}

type Credentials struct {
	UserId   int    `json:"id"`
	Password string `json:"password"`
}

type ReqUser struct {
	Id       int    `json:"id"`
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
