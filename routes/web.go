package routes

import (
	"icepeak/app/controllers"
	"icepeak/core/routing"
	"net/http"
)

// RegisterWebRoutes registers web routes
func RegisterWebRoutes(router *routing.Router, viewRoot string, validationMiddleware func(http.Handler) http.Handler) {
	// Home route - No input validation middleware required here
	router.Get("/", controllers.HomeController(viewRoot))

	// Example of applying InputValidationMiddleware to a specific route
	router.Get("/submit", controllers.HomeController(viewRoot), validationMiddleware)
}
