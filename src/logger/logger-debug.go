package logger

import "log"

type loggerDebug struct {
	logger *log.Logger
}

func (l *loggerDebug) Printf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerDebug) Println(args...interface{}) {
	l.logger.Println(args...)
}