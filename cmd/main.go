package main

import (
	"net/http"

	"golang-crud-server/db"
	"golang-crud-server/logger"
	mw "golang-crud-server/middleware"
	"golang-crud-server/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {
	
	//initialize database and get router
	err := db.DBinit()
	if err != nil {
		logger.ErrorLog.Println(err)
		return
	}
	router := routes.Router()

	//CORS handling
	h := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	m := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	o := handlers.AllowedOrigins([]string{"*"})

	logger.CommonLog.Println("Starting server at port 9091")
	err = http.ListenAndServe(":9091", handlers.CORS(h, m, o)(mw.Authorize(router)))
	if err!=nil{
		logger.ErrorLog.Println(err)
	}
}

