package db

import (
	"database/sql"
	
	"golang-crud-server/logger"
)

var db *sql.DB
var err error

//return pointer to db object
func GetDatabase() *sql.DB {
	return db
}

func DBinit() error {

	//open new database connection
	db, err = sql.Open("mysql", "root:Mobile@123@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil {
		logger.ErrorLog.Println(err)
		return err
	}
	
	//check if db is connected
	err = db.Ping()
	if err != nil {
		logger.ErrorLog.Println(err)
		return err
	}
	logger.CommonLog.Println("Database Connected")
	return nil
}
