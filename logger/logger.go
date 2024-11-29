package logger

import (
	"log"
	"os"
)

var (
	Error *log.Logger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	Info  *log.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
)
