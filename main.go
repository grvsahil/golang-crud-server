package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grvsahil/projectEmployeeJS/database"
	"github.com/grvsahil/projectEmployeeJS/logger"
	"github.com/grvsahil/projectEmployeeJS/route"
)

func main() {
	database.DBinit()
	router := route.Router()
	logger.CommonLog.Println("Starting server at port 9091")
	logger.ErrorLog.Println(http.ListenAndServe(":9091",router))
}