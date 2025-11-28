package dependency_test

import (
	"errors"
	"testing"

	"github.com/securehaven/dependency"
	"github.com/stretchr/testify/assert"
)

type MyDependency struct {
	Count int
}

var MyDependencyName = dependency.Name(new(MyDependency))
var MyDependencyError = errors.New("something went wrong")

func AddMyDependency() dependency.DependencyFunc {
	return func() (string, dependency.FactoryFunc) {
		return MyDependencyName, func(c *dependency.Container) (any, error) {
			return &MyDependency{}, nil
		}
	}
}

func AddMyDependencyError() dependency.DependencyFunc {
	return func() (string, dependency.FactoryFunc) {
		return MyDependencyName, func(c *dependency.Container) (any, error) {
			return &MyDependency{}, MyDependencyError
		}
	}
}

func TestContainerResolve_Success(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependency())
	myDep, err := container.Resolve(MyDependencyName)

	assert.NoError(err)
	assert.IsType(new(MyDependency), myDep)
}

func TestContainerResolveMultiple_Success(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependency())

	for range 3 {
		myDep, err := container.Resolve(MyDependencyName)

		assert.NoError(err)
		assert.IsType(new(MyDependency), myDep)

		myDep.(*MyDependency).Count++

		t.Logf("%#v", myDep)
	}
}

func TestContainerResolveMissingDependency_Error(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer()
	_, err := container.Resolve(MyDependencyName)

	assert.ErrorIs(err, dependency.ErrFactoryNotFound)
}

func TestContainerResolveDependencyWithError_Error(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependencyError())
	_, err := container.Resolve(MyDependencyName)

	assert.ErrorIs(err, MyDependencyError)
}
