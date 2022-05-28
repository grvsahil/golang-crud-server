package routes

import (
	"golang-crud-server/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controller.List).Methods("GET") //read
	r.HandleFunc("/user/{id}", controller.Update).Methods("PATCH") //update 
	r.HandleFunc("/user/{id}", controller.Delete).Methods("DELETE") //delete
	r.HandleFunc("/user", controller.Register).Methods("POST") //create
	r.HandleFunc("/login", controller.Login).Methods("POST") //login

	return r
}
