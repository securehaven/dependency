# Dependency: A straightforward DI container for Go

Dependency is a lightweigt and easy-to-use dependency injection library for Go, featuring named dependencies and type-safe resolution.

## Getting started

```sh
go get -u github.com/securehaven/dependency
```

## Examples

```go
import (
	"github.com/securehaven/dependency"
)

type MyDependency struct{}

// The name is used to resolve the dependency later.
// Using the dependency.Name() helper allows us to resolve
// a dependency via a generic type e.g. dependency.Resolve[*MyDependency]().
var MyDependencyName = dependency.Name(new(MyDependency))

func AddMyDependency() dependency.DependencyFunc {
	return func() (string, dependency.FactoryFunc) {
		return MyDependencyName, func(c *dependency.Container) (any, error) {
			return &MyDependency{}, nil
		}
	}
}
```

### Using the standard container

```go
// Register a single or multiple dependencies.
dependency.Register(
	AddMyDependency(),
)

// Resolve a dependency by accessing the standard container directly.
myDependency, err := dependency.StdContainer.Resolve(MyDependencyName)
	
// An error is returned when the dependency cannot be resolved.
// This is the case when either the dependency nor the factory can be found by name.
// Usually this happens when the dependency is not registered or
// the dependency name is not equal to the returned name from the DependencyFunc.
if err != nil {
	log.Fatal(err)
}

// Resolving a dependency via a generic type.
myDependency, err := dependency.Resolve[*MyDependency]()

// Panics instead of returning the error.
myDependency := dependency.MustResolve[*MyDependency]()
```

### Using a custom container

```go
// Create a new container instance.
container := dependency.NewContainer(
	AddMyDependency(),
)

// Register additional dependencies.
container.Register(
	AddMyDependency(),
)

myDependency, err := container.Resolve(MyDependencyName)

// An error is returned when the dependency cannot be resolved.
// This is the case when either the dependency nor the factory can be found by name.
// Usually this happens when the dependency is not registered or
// the dependency name is not equal to the returned name from the DependencyFunc.
if err != nil {
	log.Fatal(err)
}

// Resolving a dependency via a generic type using a custom container as resolver.
myDependeny, err := dependency.ResolveWithResolver[*MyDependency](container)

// Panics instead of returning the error.
myDependeny := dependency.MustResolveWithResolver[*MyDependency](container)
```
