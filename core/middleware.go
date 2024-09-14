package core

import (
	"fmt"
	"net/http"
	"strings"
)

// CORSOptions defines the configuration for CORS handling.
type CORSOptions struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

// CORSMiddleware handles Cross-Origin Resource Sharing.
func CORSMiddleware(options CORSOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" && isOriginAllowed(origin, options.AllowedOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(options.AllowedMethods, ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(options.AllowedHeaders, ", "))

				// Handle preflight OPTIONS request
				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Helper function to check if an origin is allowed.
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if allowed == "*" || strings.EqualFold(origin, allowed) {
			return true
		}
	}
	return false
}

// InputValidationMiddleware validates the incoming request data.
func InputValidationMiddleware(requiredFields []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check for required fields
			missingFields := []string{}
			for _, field := range requiredFields {
				if r.FormValue(field) == "" {
					missingFields = append(missingFields, field)
				}
			}

			// If there are missing fields, return an error response
			if len(missingFields) > 0 {
				http.Error(w, "Missing required fields: "+strings.Join(missingFields, ", "), http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequestLoggingMiddleware logs details about the incoming request.
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[REQUEST] %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
