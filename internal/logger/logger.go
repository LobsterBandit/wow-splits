package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func CreateGlobalLogger(debug bool) {
	if debug {
		Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	} else {
		Logger = log.New(os.Stdout, "", 0)
	}
}
