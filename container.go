package dependency

import (
	"errors"
	"fmt"
	"sync"
)

type FactoryFunc func(c *Container) (any, error)

type DependencyFunc func() (string, FactoryFunc)

type Registar interface {
	Register(deps ...DependencyFunc)
}

type Resolver interface {
	Resolve(key string) (any, error)
}

type Container struct {
	mutex        sync.Mutex
	factories    map[string]FactoryFunc
	dependencies map[string]any
}

var ErrFactoryNotFound = errors.New("factory not found")

func NewContainer(deps ...DependencyFunc) *Container {
	factories := make(map[string]FactoryFunc, len(deps))

	for _, dep := range deps {
		name, factory := dep()
		factories[name] = factory
	}

	return &Container{
		factories:    factories,
		dependencies: make(map[string]any),
	}
}

func (c *Container) Register(deps ...DependencyFunc) {
	for _, dep := range deps {
		name, factory := dep()

		c.mutex.Lock()
		c.factories[name] = factory
		c.mutex.Unlock()
	}
}

func (c *Container) Resolve(key string) (any, error) {
	c.mutex.Lock()
	dependency, ok := c.dependencies[key]
	c.mutex.Unlock()

	if ok {
		return dependency, nil
	}

	c.mutex.Lock()
	factory, ok := c.factories[key]
	c.mutex.Unlock()

	if !ok {
		return nil, fmt.Errorf("%w for %s", ErrFactoryNotFound, key)
	}

	dep, err := factory(c)

	if err != nil {
		return nil, fmt.Errorf("error from factory for %s: %w", key, err)
	}

	c.mutex.Lock()
	c.dependencies[key] = dep
	c.mutex.Unlock()

	return dep, nil
}
