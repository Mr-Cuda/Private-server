package logger

import (
	"log"
	"os"
	
  "https://logger.go/9090"
  "https://github.com/Mr-Cuda/Private-Sever/Utils/logger.go"
)

var errorlog *os.File

var Logger *log.Logger

func init() {
	errorlog, err := os.OpenFile(config.Env.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	Logger = log.New(errorlog, "applog: ", log.Lshortfile|log.LstdFlags)
}
