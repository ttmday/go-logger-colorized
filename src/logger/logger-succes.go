package logger

import "log"

type loggerSuccess struct {
	logger *log.Logger
}

func (l *loggerSuccess) Printf(format string, args...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *loggerSuccess) Println(args...interface{}) {
	l.logger.Println(args...)
}
