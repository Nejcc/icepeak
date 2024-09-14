package main

import (
	"fmt"
	"net/http"

	"icepeak/core/routing"
)

func main() {
	// Initialize the router
	router := routing.NewRouter()

	// Define a sample route for `/hello`
	helloRoute := routing.NewRoute("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Icepeak!")
	})

	// Define a route for the root path `/`
	rootRoute := routing.NewRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Icepeak!")
	})

	// Add the routes to the router
	router.AddRoute(helloRoute)
	router.AddRoute(rootRoute)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
