package routing

import (
	"net/http"
)

// Route represents a single route in the application.
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Name        string
	Middleware  []func(http.Handler) http.Handler
	Params      map[string]string // To store dynamic route parameters
}

// NewRoute creates a new route.
func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	return &Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
		Middleware:  []func(http.Handler) http.Handler{},
		Params:      make(map[string]string),
	}
}

// Helper functions for common HTTP methods:

// Get creates a new GET route.
func Get(path string, handler http.HandlerFunc) *Route {
	return NewRoute("GET", path, handler)
}

// Post creates a new POST route.
func Post(path string, handler http.HandlerFunc) *Route {
	return NewRoute("POST", path, handler)
}

// Put creates a new PUT route.
func Put(path string, handler http.HandlerFunc) *Route {
	return NewRoute("PUT", path, handler)
}

// Delete creates a new DELETE route.
func Delete(path string, handler http.HandlerFunc) *Route {
	return NewRoute("DELETE", path, handler)
}

// WithMiddleware adds middleware to the route.
func (r *Route) WithMiddleware(middleware ...func(http.Handler) http.Handler) *Route {
	r.Middleware = append(r.Middleware, middleware...)
	return r
}

// WithName assigns a name to the route.
func (r *Route) WithName(name string) *Route {
	r.Name = name
	return r
}
