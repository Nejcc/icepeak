package core

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// ServiceContainer is the core of the DI system in Icepeak
type ServiceContainer struct {
	services   map[string]interface{}        // Holds the actual service instances
	factories  map[string]func() interface{} // Holds factory functions for lazy loading services
	singletons map[string]bool               // Tracks which services are singletons
	lock       sync.RWMutex                  // Ensures thread-safe access
}

// NewServiceContainer initializes a new ServiceContainer
func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		services:   make(map[string]interface{}),
		factories:  make(map[string]func() interface{}),
		singletons: make(map[string]bool),
	}
}

// Register registers a new service with an optional singleton flag
func (sc *ServiceContainer) Register(name string, factory func() interface{}, isSingleton bool) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	sc.factories[name] = factory
	sc.singletons[name] = isSingleton
}

// Resolve resolves a service by name, with support for lazy loading and singletons
func (sc *ServiceContainer) Resolve(name string) (interface{}, error) {
	sc.lock.RLock()
	service, exists := sc.services[name]
	sc.lock.RUnlock()

	if exists {
		return service, nil // Return the already resolved service
	}

	sc.lock.RLock()
	factory, exists := sc.factories[name]
	sc.lock.RUnlock()

	if !exists {
		return nil, errors.New(fmt.Sprintf("Service '%s' not registered", name))
	}

	// Instantiate the service using its factory function
	service = factory()

	// If it's a singleton, store the instance for future use
	if sc.singletons[name] {
		sc.lock.Lock()
		sc.services[name] = service
		sc.lock.Unlock()
	}

	return service, nil
}

// AutoResolve attempts to resolve dependencies dynamically based on the type
func (sc *ServiceContainer) AutoResolve(target interface{}) error {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}

	elem := value.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("target must point to a struct")
	}

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := elem.Type().Field(i)

		// Check for an "inject" tag
		tag := fieldType.Tag.Get("inject")
		if tag == "" || !field.CanSet() {
			continue
		}

		// Resolve the service by the tag name
		service, err := sc.Resolve(tag)
		if err != nil {
			return err
		}

		// Set the field with the resolved service
		field.Set(reflect.ValueOf(service))
	}

	return nil
}

// RegisterSingleton is a helper method to register a singleton service
func (sc *ServiceContainer) RegisterSingleton(name string, factory func() interface{}) {
	sc.Register(name, factory, true)
}

// RegisterLazy registers a service that should be lazily instantiated
func (sc *ServiceContainer) RegisterLazy(name string, factory func() interface{}) {
	sc.Register(name, factory, false)
}
