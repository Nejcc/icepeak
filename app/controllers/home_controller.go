package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeController handles requests to the home page
func HomeController(viewRoot string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Load and render the home view
		tmpl, err := template.ParseFiles(filepath.Join(viewRoot, "welcome/index.html"))
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}
