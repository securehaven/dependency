package dependency

// Standard container.
// Use this if you only need a single container instance.
var StdContainer = NewContainer()

// Set a new standard container.
func SetStdContainer(c *Container) {
	StdContainer = c
}

// Registers dependencies using the standard container (StdContainer).
func Register(deps ...DependencyFunc) {
	StdContainer.Register(deps...)
}

// Resolves a dependency using the standard container (StdContainer).
// Panics on error.
//
// Deprecated: Use Must(Resolve()) instead.
func MustResolve[Dependency any]() Dependency {
	return MustResolveWithResolver[Dependency](StdContainer)
}

// Resolves a dependency using the standard container (StdContainer).
func Resolve[Dependency any]() (Dependency, error) {
	return ResolveWithResolver[Dependency](StdContainer)
}
