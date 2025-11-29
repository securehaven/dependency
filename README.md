# 🚀 Dependency: A Lightweight DI Container for Go

**Dependency** is a straightforward, lightweight, and easy-to-use **Dependency Injection (DI) container** for Go, designed for speed and **type-safe resolution** using Go Generics.

-----

## 📦 Getting Started

Install the library using the standard Go toolchain:

```sh
go get -u github.com/securehaven/dependency
```

-----

## ✨ Core Concepts & Registration

Dependencies are registered using a unique **name** and a **FactoryFunc** that defines how to create the instance.

The helper `dependency.Name(T)` is crucial as it generates a string name from the *type*, enabling the type-safe generic resolution later.

### 📝 Example: Defining and Registering a Dependency

```go
package main

import (
	"log"
	"github.com/securehaven/dependency"
)

// Define your dependency struct
type MyDependency struct{
    Value string
}

// 1. Generate a unique name for the dependency based on its type (*MyDependency).
var MyDependencyName = dependency.Name(new(MyDependency))

// 2. Define a DependencyFunc (a factory wrapper)
var MyDependency dependency.DependencyFunc = func() (string, dependency.FactoryFunc) {
    // Returns the unique name and the factory function
	return MyDependencyName, func(c *dependency.Container) (any, error) {
		// The factory creates and returns the concrete instance
        log.Println("MyDependency instance created.")
		return &MyDependency{Value: "Injected"}, nil
	}
}

func main() {
    // 3. Register the dependency with the standard container
    dependency.Register(
        MyDependency,
    )

    // ... resolution examples below
}
```

-----

## 🛠️ Resolving Dependencies

**Dependency** excels at type-safe resolution using generics, eliminating the need for boilerplate type assertion.

### 1. Type-Safe Generic Resolution (Recommended)

This is the cleanest way to resolve dependencies and relies on the type being used to generate the unique name during registration.

| Method | Description |
| :--- | :--- |
| `dependency.Resolve[T]()` | Resolves from the **standard container**. Returns `(T, error)`. |
| `dependency.MustResolve[T]()` | Resolves from the **standard container**. Returns `T` or **panics** on error. |
| `dependency.ResolveWithResolver[T](container)` | Resolves from a **custom container**. Returns `(T, error)`. |
| `dependency.MustResolveWithResolver[T](container)` | Resolves from a **custom container**. Returns `T` or **panics**. |

```go
// Resolving from the standard container
myDependency, err := dependency.Resolve[*MyDependency]()
if err != nil {
	log.Fatal("Failed to resolve MyDependency:", err)
}

log.Println("Resolved value:", myDependency.Value) // Output: Injected

// Panics if not found
myDependency = dependency.MustResolve[*MyDependency]()
```

### 2. Resolution by Name

You can also resolve directly using the dependency name, which requires a type assertion.

```go
// Resolve by name (less type-safe, requires assertion)
rawDep, err := dependency.StdContainer.Resolve(MyDependencyName)
if err != nil {
    log.Fatal(err)
}

myDependency, ok := rawDep.(*MyDependency)
if !ok {
    log.Fatal("Type assertion failed")
}
```

-----

## 🔗 Chained Resolution

For resolving multiple dependencies sequentially, the `Start()` and `Then()` helpers streamline the process by automatically chaining error handling.

This pattern is especially useful in initialization code where multiple dependencies must be resolved one after the other.

```go
// Assume AddFirstDependency, AddSecondDependency, etc., are registered.

// 1. Start the chain by resolving the first instance.
firstDependency := dependency.Start[*FirstDependency](dependency.StdContainer)

// 2. Chain subsequent resolutions. The previous result (firstDependency)
// is passed as an argument. The error is wrapped and propagated.
secondDependency := dependency.Then[*SecondDependency](dependency.StdContainer, firstDependency)
thirdDependency := dependency.Then[*ThirdDependency](dependency.StdContainer, secondDependency)

// 3. Only the last error needs to be checked, as it contains the full error chain.
if thirdDependency.Err != nil {
	log.Fatal("Chained resolution failed:", thirdDependency.Err)
}

// Access the resolved instances
app := NewApp(
	firstDependency.Value,
	secondDependency.Value,
	thirdDependency.Value,
)
```

-----

## 🧺 Using a Custom Container

While the `dependency.StdContainer` is convenient, you can create and manage isolated container instances.

```go
// 1. Create a new container instance and register initial dependencies.
container := dependency.NewContainer(
	MyDependency,
)

// 2. Register additional dependencies later.
container.Register(
	AnotherDependency,
)

// 3. Resolve using the custom container via the specialized generic helpers
myDependency, err := dependency.ResolveWithResolver[*MyDependency](container)
```
