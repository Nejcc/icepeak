package routing

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// Router represents the router that holds all registered routes.
type Router struct {
	routes        []*Route
	errorHandlers map[int]http.HandlerFunc // Custom error handlers
	logger        Logger                   // Logger for centralized logging
}

// Logger interface for centralized logging
type Logger interface {
	Error(msg string)
	Info(msg string)
}

// NewRouter initializes a new router.
func NewRouter() *Router {
	return &Router{
		routes:        []*Route{},
		errorHandlers: make(map[int]http.HandlerFunc),
		logger:        &DefaultLogger{}, // Using a default logger
	}
}

// AddRoute registers a new route to the router.
func (r *Router) AddRoute(route *Route) {
	r.routes = append(r.routes, route)
}

// Group defines a group of routes with a common prefix or middleware.
func (r *Router) Group(prefix string, middleware ...func(http.Handler) http.Handler) *Router {
	groupRouter := NewRouter()

	for _, route := range r.routes {
		route.Path = prefix + route.Path
		for _, mw := range middleware {
			route.WithMiddleware(mw)
		}
		groupRouter.AddRoute(route)
	}

	return groupRouter
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			env := os.Getenv("ENVIRONMENT")
			if env == "development" {
				// Detailed error message for development
				stackTrace := make([]byte, 1024)
				runtime.Stack(stackTrace, false)
				r.logger.Error(fmt.Sprintf("Internal Server Error: %v\nStack Trace:\n%s", err, stackTrace))
				r.handleError(w, req, http.StatusInternalServerError)
			} else {
				// Generic error message for production
				r.logger.Error(fmt.Sprintf("Internal Server Error: %v", err))
				r.handleError(w, req, http.StatusInternalServerError)
			}
		}
	}()

	for _, route := range r.routes {
		if route.Method == req.Method && r.matchPath(route, req.URL.Path) {
			// Wrap the route handler with context setting logic
			finalHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// Pass parameters to the handler using the context
				ctx := req.Context()
				for key, value := range route.Params {
					ctx = context.WithValue(ctx, key, value)
				}
				req = req.WithContext(ctx)
				route.HandlerFunc(w, req)
			})

			// Apply middleware
			var handler http.Handler = finalHandler
			for _, mw := range route.Middleware {
				handler = mw(handler)
			}

			// Serve the request with the final handler
			handler.ServeHTTP(w, req)
			return
		}
	}

	r.handleError(w, req, http.StatusNotFound)
}

// matchPath checks if the request path matches the route path and extracts parameters.
func (r *Router) matchPath(route *Route, requestPath string) bool {
	routeParts := strings.Split(route.Path, "/")
	requestParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(requestParts) {
		return false
	}

	// Clear any previously stored parameters
	route.Params = make(map[string]string)

	for i, part := range routeParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			// Extract dynamic parameter
			paramName := strings.Trim(part, "{}")
			route.Params[paramName] = requestParts[i]
		} else if part != requestParts[i] {
			return false
		}
	}
	return true
}

// RegisterErrorHandler allows registration of custom error handlers.
func (r *Router) RegisterErrorHandler(statusCode int, handler http.HandlerFunc) {
	r.errorHandlers[statusCode] = handler
}

// handleError handles HTTP errors using custom or default error handlers.
func (r *Router) handleError(w http.ResponseWriter, req *http.Request, statusCode int) {
	if handler, exists := r.errorHandlers[statusCode]; exists {
		r.logger.Info(fmt.Sprintf("Handling error %d for path %s", statusCode, req.URL.Path))
		handler(w, req)
	} else {
		r.logger.Error(fmt.Sprintf("Unhandled error %d for path %s", statusCode, req.URL.Path))
		http.Error(w, http.StatusText(statusCode), statusCode)
	}
}

// DefaultLogger is a basic implementation of the Logger interface.
type DefaultLogger struct{}

// Error logs an error message.
func (l *DefaultLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}

// Info logs an informational message.
func (l *DefaultLogger) Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}
