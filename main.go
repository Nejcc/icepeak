package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"

	"icepeak/core/routing"

	"github.com/joho/godotenv"
)

// Config struct to hold the configuration values
type Config struct {
	ViewRoot string `yaml:"VIEW_ROOT"`
}

var config Config

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Load the view configuration
	err = loadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	// Initialize the router
	router := routing.NewRouter()

	// Root route to render the HTML view
	rootRoute := routing.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "welcome/index.html")
	})

	// Register routes
	router.AddRoute(rootRoute)

	// Register custom error handlers
	router.RegisterErrorHandler(404, func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "errors/404.html")
	})

	router.RegisterErrorHandler(500, func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv("ENVIRONMENT")
		if env == "development" {
			// Detailed error messages for development
			stackTrace := make([]byte, 1024)
			runtime.Stack(stackTrace, false)
			renderTemplate(w, "errors/debug.html", string(stackTrace))
		} else {
			// Generic error message for production
			renderTemplate(w, "errors/500.html")
		}
	})

	router.RegisterErrorHandler(403, func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "errors/403.html")
	})

	// Fallback route for unmatched paths
	fallbackRoute := routing.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "errors/404.html")
	})
	router.AddRoute(fallbackRoute)

	// Under construction route example
	underConstructionRoute := routing.Get("/under-construction", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "errors/under_construction.html")
	})
	router.AddRoute(underConstructionRoute)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

// loadConfig reads the configuration from the YAML file
func loadConfig() error {
	data, err := ioutil.ReadFile("config/view.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &config)
}

// renderTemplate loads and renders an HTML template
func renderTemplate(w http.ResponseWriter, templateName string, data ...interface{}) {
	tmpl, err := template.ParseFiles(filepath.Join(config.ViewRoot, templateName))
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	if len(data) > 0 {
		tmpl.Execute(w, data[0])
	} else {
		tmpl.Execute(w, nil)
	}
}
