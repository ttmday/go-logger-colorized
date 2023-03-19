package logger

import "log"

type loggerInfo struct {
	logger *log.Logger
}

func (l *loggerInfo) Printf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerInfo) Println(args...interface{}) {
	l.logger.Println(args...)
}
