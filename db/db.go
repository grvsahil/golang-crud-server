package db

import (
	"database/sql"
	
	"golang-crud-server/logger"
)

var Db *sql.DB
var err error

//return pointer to db object
func GetDatabase() *sql.DB {
	return Db
}

func DBinit() {

	//open new database connection
	Db, err = sql.Open("mysql", "root:Mobile@123@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil {
		logger.ErrorLog.Println(err)
		return
	}
	
	//check if db is connected
	err = Db.Ping()
	if err != nil {
		logger.ErrorLog.Println(err)
		return
	}
	logger.CommonLog.Println("Database Connected")
}
