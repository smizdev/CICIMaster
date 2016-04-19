package ioc

import (
	"fmt"
)

// Container
type Container struct {
	// factories
	factories map[string]func() interface{}

	// singletons
	singletons map[string]interface{}
}

// NewContainer
func NewContainer() *Container {
	return &Container{
		factories:  make(map[string]func() interface{}),
		singletons: make(map[string]interface{}),
	}
}

// Register
func (t *Container) Register(name string, factory func() interface{}) {
	t.factories[name] = factory
}

// Factory
func (t *Container) Factory(name string) func() interface{} {
	return t.factories[name]
}

// Instance
func (t *Container) Instance(name string) interface{} {
	factory, ok := t.factories[name]
	if !ok {
		panic(fmt.Sprintf("No such factory registered with name %s", name))
	}
	return factory()
}

// Singleton
func (t *Container) Singleton(name string) interface{} {
	singleton, ok := t.singletons[name]
	if !ok {
		t.singletons[name] = t.factories[name]()
		return t.singletons[name]
	}
	return singleton
}
