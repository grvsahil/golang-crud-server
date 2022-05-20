package route

import (
	"github.com/gorilla/mux"
	"golang-crud-server/controller"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user", controller.CORS((controller.Authorize(controller.GetUser)))).Methods("GET")
	r.HandleFunc("/user/{id}", controller.CORS(controller.Authorize(controller.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", controller.CORS(controller.Authorize(controller.DeleteUser))).Methods("DELETE")
	r.HandleFunc("/user", controller.CORS(controller.CreateUser)).Methods("POST")
	r.HandleFunc("/login", controller.CORS(controller.LoginUser)).Methods("POST")
	r.HandleFunc("/logout", controller.CORS(controller.LogoutUser)).Methods("GET")
	r.Handle("/favicon.ico", http.NotFoundHandler())

	return r
}
