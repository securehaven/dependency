package dependency

import (
	"fmt"
	"reflect"
)

// Get the type name of a dependency.
func Name(dep any) string {
	return reflect.TypeOf(dep).String()
}

// Force a dependency without an error.
// Panics on error.
func Must[D any](dep D, err error) D {
	if err != nil {
		panic(err)
	}

	return dep
}

type Resolved[V any] struct {
	Value V
	Err   error
}

// Implements the standard error interface.
func (r Resolved[V]) Error() string {
	return r.Err.Error()
}

// Resolves a single instance and wraps it with the Resolved type including the error.
func Start[V any](r Resolver) Resolved[V] {
	dep, err := ResolveWithResolver[V](r)

	if err != nil {
		err = fmt.Errorf("failed to resolve initial dependency (%T): %w", dep, err)
	}

	return Resolved[V]{
		Value: dep,
		Err:   err,
	}
}

// Resolves another instance after checking the error of the previous one.
func Then[O, I any](resolver Resolver, input Resolved[I]) Resolved[O] {
	if input.Err != nil {
		var zero O
		return Resolved[O]{Err: input.Err, Value: zero}
	}

	dep, err := ResolveWithResolver[O](resolver)

	if err != nil {
		err = fmt.Errorf("failed to resolve subsequent dependency (%T): %w", dep, err)
	}

	return Resolved[O]{
		Value: dep,
		Err:   err,
	}
}
