package routing

import (
	"fmt"
	"net/http"
)

// Router represents the router that holds all registered routes.
type Router struct {
	routes        []*Route
	errorHandlers map[int]http.HandlerFunc // Custom error handlers
	parentRouter  *Router                  // Reference to the parent router, if any
	prefix        string                   // Prefix for routes in this group
	middleware    []func(http.Handler) http.Handler
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
	return &Router{
		parentRouter:  r,
		prefix:        r.prefix + prefix, // Carry forward any existing prefix
		middleware:    append(r.middleware, middleware...),
		errorHandlers: r.errorHandlers, // Use the same error handlers as the parent
	}
}

// AddRoute registers a new route to the router or its parent.
func (r *Router) AddRoute(route *Route) {
	// Apply the group's prefix and middleware
	route.Path = r.prefix + route.Path
	route.Middleware = append(r.middleware, route.Middleware...)

	// If this router has a parent, register the route with the parent
	if r.parentRouter != nil {
		r.parentRouter.AddRoute(route)
	} else {
		// Otherwise, register the route with this router
		r.routes = append(r.routes, route)
	}
}

// Register HTTP methods with optional middleware
func (r *Router) Get(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("GET", path, handler, middleware...))
}

func (r *Router) Post(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("POST", path, handler, middleware...))
}

func (r *Router) Put(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("PUT", path, handler, middleware...))
}

func (r *Router) Patch(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("PATCH", path, handler, middleware...))
}

func (r *Router) Delete(path string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
	r.AddRoute(NewRoute("DELETE", path, handler, middleware...))
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	errorHandler := NewErrorHandler() // Initialize the centralized error handler

	for _, route := range r.routes {
		if route.Method == req.Method && route.Pattern.MatchString(req.URL.Path) {
			// Extract parameters from the URL path
			matches := route.Pattern.FindStringSubmatch(req.URL.Path)
			if matches != nil {
				for i, name := range route.Pattern.SubexpNames() {
					if i > 0 && name != "" {
						route.Params[name] = matches[i]
					}
				}

				handler := http.Handler(http.HandlerFunc(route.HandlerFunc))
				for _, mw := range route.Middleware {
					handler = mw(handler)
				}

				// Error handling for route-specific errors
				defer func() {
					if err := recover(); err != nil {
						if route.ErrorHandler != nil {
							route.ErrorHandler(w, req)
						} else {
							errorHandler.HandleError(w, req, http.StatusInternalServerError, fmt.Errorf("%v", err))
						}
					}
				}()

				handler.ServeHTTP(w, req)
				return
			}
		}
	}
	errorHandler.HandleError(w, req, http.StatusNotFound, fmt.Errorf("Page not found"))
}

// handleError handles HTTP errors using custom or default error handlers.
func (r *Router) handleError(w http.ResponseWriter, req *http.Request, statusCode int) {
	if handler, exists := r.errorHandlers[statusCode]; exists {
		handler(w, req)
	} else {
		http.Error(w, http.StatusText(statusCode), statusCode)
	}
}
