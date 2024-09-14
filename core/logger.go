package core

import "fmt"

// DefaultLogger is a basic implementation of a logger service.
type DefaultLogger struct{}

// Error logs an error message.
func (l *DefaultLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}

// Info logs an informational message.
func (l *DefaultLogger) Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

// Debug logs a debug message.
func (l *DefaultLogger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}
