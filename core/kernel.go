package core

import (
	"fmt"
	"net/http"
	"os"

	"icepeak/core/routing"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// Kernel is the core of the Icepeak framework
type Kernel struct {
	Router     *routing.Router
	Middleware []func(http.Handler) http.Handler
	Config     map[string]interface{}
	Services   *ServiceContainer
}

var kernelInstance *Kernel

// NewKernel creates a new instance of the Kernel
func NewKernel() *Kernel {
	if kernelInstance == nil {
		kernelInstance = &Kernel{
			Router:     routing.NewRouter(),
			Middleware: []func(http.Handler) http.Handler{},
			Config:     make(map[string]interface{}),
			Services:   NewServiceContainer(),
		}
		kernelInstance.loadEnvironment()
		kernelInstance.loadConfiguration()
		kernelInstance.registerDefaultServices()
	}
	return kernelInstance
}

// GetKernel returns the current kernel instance
func GetKernel() *Kernel {
	return kernelInstance
}

// loadEnvironment loads environment variables from the .env file
func (k *Kernel) loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// loadConfiguration loads the configuration from YAML files
func (k *Kernel) loadConfiguration() {
	data, err := os.ReadFile("config/view.yaml")
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}
	err = yaml.Unmarshal(data, &k.Config)
	if err != nil {
		fmt.Printf("Error parsing configuration: %v\n", err)
	}
}

// RegisterMiddleware registers middleware to be applied to all routes
func (k *Kernel) RegisterMiddleware(middleware func(http.Handler) http.Handler) {
	k.Middleware = append(k.Middleware, middleware)
}

// registerDefaultServices registers default services in the service container
func (k *Kernel) registerDefaultServices() {
	k.Services.Register("logger", func() interface{} {
		return &DefaultLogger{}
	}, true) // Registering logger as a singleton
}

// HandleRequest manages the request lifecycle
func (k *Kernel) HandleRequest(w http.ResponseWriter, req *http.Request) {
	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		k.Router.ServeHTTP(w, req)
	}))

	// Apply all registered middleware
	for _, mw := range k.Middleware {
		handler = mw(handler)
	}

	handler.ServeHTTP(w, req)
}

// StartServer starts the HTTP server
func (k *Kernel) StartServer(address string) {
	fmt.Printf("Server running at %s\n", address)
	http.ListenAndServe(address, http.HandlerFunc(k.HandleRequest))
}
