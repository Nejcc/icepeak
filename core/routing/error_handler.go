package routing

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime/debug"
)

// ErrorHandler handles HTTP errors with custom views or messages.
type ErrorHandler struct {
	isDevMode bool
	viewPath  string
}

// NewErrorHandler initializes a new ErrorHandler.
func NewErrorHandler() *ErrorHandler {
	// Check if the environment is development or production.
	isDevMode := os.Getenv("APP_ENV") == "development"

	// Define the path where custom error views are located.
	viewPath := "resources/views/errors/"

	return &ErrorHandler{
		isDevMode: isDevMode,
		viewPath:  viewPath,
	}
}

// HandleError displays the appropriate error page or message.
func (eh *ErrorHandler) HandleError(w http.ResponseWriter, req *http.Request, statusCode int, err error) {
	if eh.isDevMode {
		// In development mode, provide a detailed error message with a stack trace.
		http.Error(w, fmt.Sprintf("Error %d: %v\n\n%s", statusCode, err, debug.Stack()), statusCode)
	} else {
		// In production mode, serve a custom error page.
		tmplPath := fmt.Sprintf("%s%d.html", eh.viewPath, statusCode)
		tmpl, tmplErr := template.ParseFiles(tmplPath)
		if tmplErr != nil {
			// Fallback if the custom error page is not found.
			http.Error(w, http.StatusText(statusCode), statusCode)
			return
		}
		tmpl.Execute(w, nil)
	}
}
