package main

import (
	"fmt"
	"net/http"

	"icepeak/core"
)

func main() {
	// Create a new kernel instance
	kernel := core.NewKernel()

	// Register middleware
	kernel.RegisterMiddleware(loggingMiddleware)

	// Start the server
	kernel.StartServer(":8080")
}

// Example middleware for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
