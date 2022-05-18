package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grvsahil/projectEmployeeJS/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", controller.CORS((controller.Authorize(controller.GetUser)))).Methods("GET")
	router.HandleFunc("/user", controller.CORS(controller.CreateUser)).Methods("POST")
	router.HandleFunc("/user/{id}", controller.CORS(controller.UpdateUser)).Methods("PATCH")
	router.HandleFunc("/user/{id}", controller.CORS(controller.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/login", controller.CORS(controller.LoginUser)).Methods("POST")
	router.HandleFunc("/logout", controller.CORS(controller.LogoutUser)).Methods("GET")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	return router
}
