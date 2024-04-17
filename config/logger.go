package config

import (
	"log"
)

type Logger struct {
	info  *log.Logger
	debug *log.Logger
	fatal *log.Logger
	error *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info:  log.New(log.Writer(), "INFO: ", log.LstdFlags),
		fatal: log.New(log.Writer(), "FATAL: ", log.LstdFlags),
		debug: log.New(log.Writer(), "DEBUG: ", log.LstdFlags),
		error: log.New(log.Writer(), "ERROR: ", log.LstdFlags),
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.fatal.Println(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.fatal.Fatalf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
