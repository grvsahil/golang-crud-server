package model

type User struct {
	ID        int64
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
	Email     string `json:"email"`
	DOB       string `json:"dob"`
	Password  string `json:"password,omitempty"`
}
