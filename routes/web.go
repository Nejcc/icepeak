package routes

import (
	"icepeak/app/controllers"
	"icepeak/core/routing"
)

// RegisterWebRoutes registers web routes
func RegisterWebRoutes(router *routing.Router, viewRoot string) {
	// Home route
	router.AddRoute(routing.Get("/", controllers.HomeController(viewRoot)))

	// More web routes can be added here...
}
