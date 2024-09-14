package routing

import (
	"net/http"
)

// Router represents the router that holds all registered routes.
type Router struct {
	routes        []*Route
	errorHandlers map[int]http.HandlerFunc // Custom error handlers
}

// NewRouter initializes a new router.
func NewRouter() *Router {
	return &Router{
		routes:        []*Route{},
		errorHandlers: make(map[int]http.HandlerFunc),
	}
}

// Group creates a new route group with a common prefix and middleware.
func (r *Router) Group(prefix string, middleware ...func(http.Handler) http.Handler) *Router {
	groupRouter := &Router{
		routes:        []*Route{},
		errorHandlers: r.errorHandlers,
	}

	// Update the group router to use the same route slice
	r.routes = append(r.routes, groupRouter.routes...)

	return &Router{
		routes:        groupRouter.routes,
		errorHandlers: groupRouter.errorHandlers,
	}
}

// Get registers a GET route with optional middleware.
func (r *Router) Get(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("GET", path, handler, middleware...))
}

// Post registers a POST route with optional middleware.
func (r *Router) Post(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("POST", path, handler, middleware...))
}

// Put registers a PUT route with optional middleware.
func (r *Router) Put(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("PUT", path, handler, middleware...))
}

// Patch registers a PATCH route with optional middleware.
func (r *Router) Patch(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("PATCH", path, handler, middleware...))
}

// Delete registers a DELETE route with optional middleware.
func (r *Router) Delete(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("DELETE", path, handler, middleware...))
}

// AddRoute registers a new route to the router.
func (r *Router) AddRoute(route *Route) {
	r.routes = append(r.routes, route)
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			handler := http.Handler(http.HandlerFunc(route.HandlerFunc))

			// Apply route-specific middleware
			for _, mw := range route.Middleware {
				handler = mw(handler)
			}

			handler.ServeHTTP(w, req)
			return
		}
	}
	r.handleError(w, req, http.StatusNotFound)
}

// handleError handles HTTP errors using custom or default error handlers.
func (r *Router) handleError(w http.ResponseWriter, req *http.Request, statusCode int) {
	if handler, exists := r.errorHandlers[statusCode]; exists {
		handler(w, req)
	} else {
		http.Error(w, http.StatusText(statusCode), statusCode)
	}
}
