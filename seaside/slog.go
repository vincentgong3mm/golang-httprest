package main

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

var sl Logger

func NewSlog() {
	sl.Info = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	sl.Error = log.New(os.Stdout, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
}
