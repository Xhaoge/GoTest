package logger


import (
	"log"
	"os"
)


type mylog struct{
	Info *log.logger
	Warn *log.logger
	Error *log.logger
}



func init(){
	logFilePath := "info_log.log"
	logfile,err := os.Create(logFilePath)
	if err != nil{
		log.Fatalln("open log file failed")
	}
	
}