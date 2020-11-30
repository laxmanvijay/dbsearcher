package utils

import (
	"log"
	"os"
)

// CustomLoggerMethods provides logging methods
type CustomLoggerMethods interface {
	init()
	Error(err interface{})
}

// Logger defines a custom logger
type Logger struct {
	OutFile string
	Logger  *log.Logger
}

// Init sets a out file
func (logger *Logger) init() {
	outfile, _ := os.Create("dbsearcher.log")
	logger.Logger = log.New(outfile, "", 0)
}

func (logger *Logger) Error(err interface{}) {
	logger.Logger.Println(err)
}

// GetLogger provides a logger
func GetLogger() CustomLoggerMethods {
	var logger Logger
	logger.init()
	return &logger
}
