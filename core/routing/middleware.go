package routing

import "net/http"

// Middleware is a function that wraps around an http.Handler to perform additional processing.
type Middleware func(http.Handler) http.Handler

// ApplyMiddleware applies middleware to a given handler.
func ApplyMiddleware(handler http.Handler, middleware ...Middleware) http.Handler {
	for _, mw := range middleware {
		handler = mw(handler)
	}
	return handler
}
