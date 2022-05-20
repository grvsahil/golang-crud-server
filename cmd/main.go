package main

import (
	"net/http"

	"ggolang-crud-server/route"
	"golang-crud-server/database"
	"golang-crud-server/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {
	database.DBinit()

	router := route.Router()

	logger.CommonLog.Println("Starting server at port 9091")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	logger.ErrorLog.Println(http.ListenAndServe(":9091", handlers.CORS(headers, methods, origins)(router)))
}
