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
}

// NewRoute creates a new route.
func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	return &Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
		Middleware:  []func(http.Handler) http.Handler{},
	}
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
