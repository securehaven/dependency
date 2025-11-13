package dependency

var StdContainer = NewContainer()

func SetStdContainer(c *Container) {
	StdContainer = c
}

func Register(deps ...DependencyFunc) {
	StdContainer.Register(deps...)
}

func MustResolve[Dependency any]() Dependency {
	return MustResolveWithResolver[Dependency](StdContainer)
}

func Resolve[Dependency any]() (Dependency, error) {
	return ResolveWithResolver[Dependency](StdContainer)
}
