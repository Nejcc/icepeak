package core

import (
	"fmt"
	"log"
	"os"
)

// Logger interface defines methods for different log levels.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// DefaultLogger provides a basic logger implementation.
type DefaultLogger struct {
	logger *log.Logger
	level  string
}

// NewDefaultLogger creates a new instance of DefaultLogger with a given log level and output.
func NewDefaultLogger(level string, output string) *DefaultLogger {
	var out *os.File
	if output == "file" {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Could not open log file: %v", err)
			out = os.Stdout
		} else {
			out = file
		}
	} else {
		out = os.Stdout
	}
	return &DefaultLogger{
		logger: log.New(out, "", log.LstdFlags),
		level:  level,
	}
}

// Debug logs a debug message.
func (l *DefaultLogger) Debug(msg string) {
	if l.level == "DEBUG" {
		l.logger.Println("[DEBUG]", msg)
	}
}

// Info logs an informational message.
func (l *DefaultLogger) Info(msg string) {
	l.logger.Println("[INFO]", msg)
}

// Warn logs a warning message.
func (l *DefaultLogger) Warn(msg string) {
	l.logger.Println("[WARN]", msg)
}

// Error logs an error message.
func (l *DefaultLogger) Error(msg string) {
	l.logger.Println("[ERROR]", msg)
}
