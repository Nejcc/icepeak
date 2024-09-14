package main

import (
	"fmt"
	"net/http"

	"icepeak/core/routing"
)

func main() {
	// Initialize the router
	router := routing.NewRouter()

	// Define routes using helper functions for different HTTP methods

	// GET route
	helloRoute := routing.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Icepeak!")
	})

	// POST route
	createUserRoute := routing.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User created!")
	})

	// PUT route
	updateUserRoute := routing.Put("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(string)
		fmt.Fprintf(w, "User ID %s updated!", id)
	})

	// DELETE route
	deleteUserRoute := routing.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(string)
		fmt.Fprintf(w, "User ID %s deleted!", id)
	})

	// Root route
	rootRoute := routing.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Icepeak!")
	})

	// Register the routes with the router
	router.AddRoute(helloRoute)
	router.AddRoute(createUserRoute)
	router.AddRoute(updateUserRoute)
	router.AddRoute(deleteUserRoute)
	router.AddRoute(rootRoute)

	// Register custom error handlers
	router.RegisterErrorHandler(404, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Custom 404: Page not found!", http.StatusNotFound)
	})

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

// loggingMiddleware is a sample middleware for logging requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
