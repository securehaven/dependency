package dependency_test

import (
	"testing"

	"github.com/securehaven/dependency"
	"github.com/stretchr/testify/assert"
)

func TestSetStdContainer_Success(t *testing.T) {
	container := dependency.NewContainer(AddMyDependency())
	dependency.SetStdContainer(container)

	assert.Equal(t, container, dependency.StdContainer)
}

func TestSetStdContainer_Error(t *testing.T) {
	container := dependency.NewContainer(AddMyDependency())

	assert.NotEqual(t, container, dependency.StdContainer)
}

func TestRegister_Success(t *testing.T) {
	assert := assert.New(t)

	dependency.Register(AddMyDependency())
	myDep, err := dependency.StdContainer.Resolve(MyDependencyName)

	assert.NoError(err)
	assert.IsType(new(MyDependency), myDep)
}

func TestResolve_Success(t *testing.T) {
	assert := assert.New(t)

	dependency.Register(AddMyDependency())
	myDep, err := dependency.Resolve[*MyDependency]()

	assert.NoError(err)
	assert.IsType(new(MyDependency), myDep)
}
