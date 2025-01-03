package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Printf("INFO: "+msg, args...)
}

func (l *Logger) Error(msg string, err error, args ...interface{}) {
	l.Printf("ERROR: "+msg+": %v", append(args, err)...)
}

func (l *Logger) Fatal(msg string, err error, args ...interface{}) {
	l.Printf("FATAL: "+msg+": %v", append(args, err)...)
	os.Exit(1)
}
