package controllers

import (
	"encoding/json"
	"net/http"
)

// UserController handles API requests for users
func UserController(w http.ResponseWriter, r *http.Request) {
	// Example JSON response
	users := []map[string]string{
		{"id": "1", "name": "John Doe"},
		{"id": "2", "name": "Jane Smith"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
