package routes

import (
	"icepeak/app/controllers"
	"icepeak/core/routing"
)

// RegisterAPIRoutes registers API routes
func RegisterAPIRoutes(router *routing.Router) {
	// Example of registering a GET route using the router instance
	router.Get("/api/users", controllers.UserController)

	// Add more API routes here as needed...
}
