package logger

import "log"

type loggerError struct {
	logger *log.Logger
}

func (l *loggerError) Printf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerError) Println(args...interface{}) {
	l.logger.Println(args...)
}


func (l *loggerError) Panicf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerError) Fatalf(format string, args...interface{}) {
	l.logger.Fatalf(format, args...)
}