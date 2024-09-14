package routing

import (
	"net/http"
	"regexp"
	"strings"
)

// Route represents an individual route with parameters, middleware, and error handling.
type Route struct {
	Method       string
	Path         string
	Pattern      *regexp.Regexp
	HandlerFunc  http.HandlerFunc
	Middleware   []func(http.Handler) http.Handler
	Params       map[string]string
	ErrorHandler http.HandlerFunc // Route-specific error handler
}

// NewRoute creates a new route instance with dynamic parameters.
func NewRoute(method, path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) *Route {
	// Convert dynamic parameters to a regex pattern
	pattern := regexp.MustCompile(`\{([a-zA-Z0-9]+)(:[^}]+)?\}`)
	regexPath := "^" + pattern.ReplaceAllStringFunc(path, func(m string) string {
		parts := strings.SplitN(m[1:len(m)-1], ":", 2)
		if len(parts) == 2 {
			return "(" + parts[1] + ")"
		}
		return "([^/]+)"
	}) + "$"

	return &Route{
		Method:      method,
		Path:        path,
		Pattern:     regexp.MustCompile(regexPath),
		HandlerFunc: handler,
		Middleware:  middleware,
		Params:      make(map[string]string),
	}
}
