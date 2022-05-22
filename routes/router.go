package routes

import (
	c "golang-crud-server/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", c.List).Methods("GET")
	r.HandleFunc("/user/{id}", c.Update).Methods("PATCH")
	r.HandleFunc("/user/{id}", c.Delete).Methods("DELETE")
	r.HandleFunc("/user", c.Register).Methods("POST")
	r.HandleFunc("/login", c.Login).Methods("POST")
	r.HandleFunc("/logout", c.Logout).Methods("GET")

	return r
}
