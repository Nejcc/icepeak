package main

import (
	"icepeak/core"
	"icepeak/routes"
)

func main() {
	// Create a new kernel instance
	kernel := core.NewKernel()

	// Register middleware globally
	kernel.RegisterMiddleware(core.RequestLoggingMiddleware)
	kernel.RegisterMiddleware(core.CORSMiddleware(core.CORSOptions{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}))

	// Register routes with selective middleware
	routes.RegisterWebRoutes(kernel.Router, kernel.Config["VIEW_ROOT"].(string), core.InputValidationMiddleware([]string{"name", "email"}))
	routes.RegisterAPIRoutes(kernel.Router)

	// Start the server
	kernel.StartServer(":8080")
}
