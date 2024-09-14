package routing

import (
	"net/http"
	"strings"
)

// Router represents the router that holds all registered routes.
type Router struct {
	routes []*Route
}

// NewRouter initializes a new router.
func NewRouter() *Router {
	return &Router{
		routes: []*Route{},
	}
}

// AddRoute registers a new route to the router.
func (r *Router) AddRoute(route *Route) {
	r.routes = append(r.routes, route)
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && matchPath(route.Path, req.URL.Path) {
			// Apply middleware
			handler := route.HandlerFunc
			for _, mw := range route.Middleware {
				handler = mw(handler)
			}
			handler.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

// matchPath checks if the request path matches the route path.
func matchPath(routePath, requestPath string) bool {
	// Basic implementation: Check if route and request paths are equal.
	return strings.TrimSuffix(routePath, "/") == strings.TrimSuffix(requestPath, "/")
}
