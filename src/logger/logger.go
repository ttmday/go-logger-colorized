package logger

import (
	"log"
	"os"
)

// COLORS //
var (
	green     = "\033[32m"
	red       = "\033[31m"
	brightRed = "\033[91m"
	gray      = "\033[90m"
	yellow    = "\033[33m"
)

func Success() *loggerSuccess {
	logger := log.New(os.Stdout, string(green)+"SUCCESS: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &loggerSuccess{
		logger: logger,
	}
}

func Info() *loggerInfo {
	logger := log.New(os.Stdout, string(gray)+"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &loggerInfo{
		logger: logger,
	}
}

func Warning() *loggerWarning {
	logger := log.New(os.Stdout, string(yellow)+"WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &loggerWarning{
		logger: logger,
	}
}

func Error() *loggerError {
	logger := log.New(os.Stdout, string(red)+"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &loggerError{
		logger: logger,
	}
}

func Debug() *loggerDebug {
	logger := log.New(os.Stdout, string(brightRed)+"DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &loggerDebug{
		logger: logger,
	}
}
