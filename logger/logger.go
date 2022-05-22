package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var CommonLog *log.Logger
var ErrorLog *log.Logger

func init() {

	//creates new log file with given permissions
	file, err := os.OpenFile("logs/"+getFilename(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	//creates new logger
	CommonLog = log.New(file, "Common Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

//eturns log file name
func getFilename() string {
	t := time.Now()
	format := "2006-01-02-15:04:05"
	return "log-"+t.Format(format) + ".log"
}
