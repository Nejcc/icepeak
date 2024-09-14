package routes

import (
	"icepeak/app/controllers"
	"icepeak/core/routing"
)

// RegisterAPIRoutes registers API routes
func RegisterAPIRoutes(router *routing.Router) {
	// API routes
	router.AddRoute(routing.Get("/api/users", controllers.UserController))

	// More API routes can be added here...
}
