package routing

import (
	"net/http"
)

// Route represents an individual route with its path, method, handler, and middleware.
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middleware  []func(http.Handler) http.Handler
}

// NewRoute creates a new route instance.
func NewRoute(method, path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) *Route {
	return &Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
		Middleware:  middleware,
	}
}
