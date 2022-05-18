package controller

import "github.com/grvsahil/projectEmployeeJS/model"

func Validate(u *model.ReqUser) bool {

	if len(u.Fname)+len(u.Lname) > 30 {
		return false
	}

	if len(u.Password) < 8 || len(u.Password) > 20 {
		return false
	}

	if len(u.Email) > 20 {
		return false
	}

	return true

}
