package main

import (
	"fmt"
	"net/http"

	"icepeak/core/routing"
)

func main() {
	// Initialize the router
	router := routing.NewRouter()

	// Define a sample route
	helloRoute := routing.NewRoute("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Icepeak!")
	})

	// Add the route to the router
	router.AddRoute(helloRoute)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
