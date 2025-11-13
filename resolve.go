package dependency

import "errors"

func MustResolveWithResolver[Dependency any](r Resolver) Dependency {
	dep, err := ResolveWithResolver[Dependency](r)

	if err != nil {
		panic(err)
	}

	return dep
}

func ResolveWithResolver[Dependency any](r Resolver) (Dependency, error) {
	var dep Dependency

	name := Name(dep)
	d, err := r.Resolve(name)

	if err != nil {
		return dep, err
	}

	dependency, ok := d.(Dependency)

	if !ok {
		return dep, errors.Join(
			ErrMissingDependency,
			ErrTypeConversion,
		)
	}

	return dependency, nil
}
