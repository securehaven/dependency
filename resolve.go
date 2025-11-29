package dependency

import (
	"errors"
	"fmt"
)

var ErrTypeConversion = errors.New("failed to convert type")

// Resolves a dependency using a resolver.
// Panics on error.
//
// Deprecated: Use Must(ResolveWithResolver()) instead.
func MustResolveWithResolver[Dependency any](r Resolver) Dependency {
	dep, err := ResolveWithResolver[Dependency](r)

	if err != nil {
		panic(err)
	}

	return dep
}

// Resolves a dependency using a resolver.
func ResolveWithResolver[Dependency any](r Resolver) (Dependency, error) {
	var dep Dependency

	name := Name(dep)
	d, err := r.Resolve(name)

	if err != nil {
		return dep, err
	}

	dependency, ok := d.(Dependency)

	if !ok {
		return dep, fmt.Errorf("%w: expected %T, received %T", ErrTypeConversion, dep, d)
	}

	return dependency, nil
}
