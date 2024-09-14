package core

// ServiceContainer manages services and their dependencies.
type ServiceContainer struct {
	services     map[string]interface{}
	singletons   map[string]interface{}
	initializers map[string]func() interface{}
}

// NewServiceContainer creates a new service container.
func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		services:     make(map[string]interface{}),
		singletons:   make(map[string]interface{}),
		initializers: make(map[string]func() interface{}),
	}
}

// Register registers a service with a given name and an initializer function.
func (c *ServiceContainer) Register(name string, initializer func() interface{}, singleton bool) {
	if singleton {
		c.singletons[name] = nil // Mark as singleton to be initialized later
	}
	c.initializers[name] = initializer
}

// Get retrieves a service by name, initializing it if necessary.
func (c *ServiceContainer) Get(name string) interface{} {
	// Check for singleton
	if singleton, exists := c.singletons[name]; exists {
		if singleton == nil {
			c.singletons[name] = c.initializers[name]()
		}
		return c.singletons[name]
	}

	// Check for regular service
	if initializer, exists := c.initializers[name]; exists {
		return initializer()
	}

	panic("Service not found: " + name)
}
