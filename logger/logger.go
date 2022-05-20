package logger

import (
	"fmt"
	"log"
	"os"
)

var CommonLog *log.Logger
var ErrorLog *log.Logger

func init() {
	openLogfile, err := os.OpenFile("/home/gaurav/go/src/github.com/grvsahil/server/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	CommonLog = log.New(openLogfile, "Common Logger:\t",   log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(openLogfile, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
