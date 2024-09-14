package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Logger interface defines methods for different log levels.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	LogRequest(r *http.Request, start time.Time)
}

// DefaultLogger provides a basic logger implementation.
type DefaultLogger struct {
	logger *log.Logger
	level  string
}

// NewDefaultLogger creates a new instance of DefaultLogger with a given log level and output.
func NewDefaultLogger(level string, output string) *DefaultLogger {
	var out io.Writer

	// Create log directory and file path
	logDir := "./storage/logs"
	logFile := filepath.Join(logDir, "app.log")

	// Ensure the log directory exists
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create log directory: %v\n", err)
		out = os.Stdout // Fallback to stdout if directory creation fails
	} else {
		// Open or create the log file
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Could not open log file: %v\n", err)
			out = os.Stdout // Fallback to stdout if file creation fails
		} else {
			// Create multi-writer to write to both file and stdout
			out = io.MultiWriter(os.Stdout, file)
		}
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

// LogRequest logs detailed information about the incoming HTTP request.
func (l *DefaultLogger) LogRequest(r *http.Request, start time.Time) {
	duration := time.Since(start)
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		clientIP = forwarded
	}

	requestSize := r.ContentLength
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	logMessage := fmt.Sprintf(
		"Client IP: %s | Method: %s | Path: %s | Request Size: %d bytes | Duration: %v | Memory Usage: %d KB",
		clientIP, r.Method, r.URL.Path, requestSize, duration, m.Alloc/1024,
	)
	l.Info(logMessage)
}
