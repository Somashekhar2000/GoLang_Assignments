package logger

import (
	"log"
)

type LoggerStore struct { //defining logger struct
	log *log.Logger
}

func New() LoggerStore { //calling new func
	return LoggerStore{
		log: log.New(log.Writer(), "soma", 56),
	}
}
