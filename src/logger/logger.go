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

func Success() *log.Logger {
	return log.New(os.Stdout, string(green)+"SUCCESS: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info() *log.Logger {
	return log.New(os.Stdout, string(gray)+"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Warning() *log.Logger {
	return log.New(os.Stdout, string(yellow)+"WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Error() *log.Logger {
	return log.New(os.Stdout, string(red)+"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug() *log.Logger {
	return log.New(os.Stdout, string(brightRed)+"DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
