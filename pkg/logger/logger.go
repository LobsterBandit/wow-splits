package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
