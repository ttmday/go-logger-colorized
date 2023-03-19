package logger

import "log"

type loggerWarning struct {
	logger *log.Logger
}

func (l *loggerWarning) Printf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerWarning) Println(args...interface{}) {
	l.logger.Println(args...)
}
