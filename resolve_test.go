package dependency_test

import (
	"testing"

	"github.com/securehaven/dependency"
	"github.com/stretchr/testify/assert"
)

func AddMyDependencyWrongType() dependency.DependencyFunc {
	return func() (string, dependency.FactoryFunc) {
		return MyDependencyName, func(c *dependency.Container) (any, error) {
			return "notTheExpectedType", nil
		}
	}
}

func TestResolveWithResolver_Success(t *testing.T) {
	c := dependency.NewContainer(AddMyDependency())
	_, err := dependency.ResolveWithResolver[*MyDependency](c)

	assert.NoError(t, err)
}

func TestResolveWithResolverMultiple_Success(t *testing.T) {
	c := dependency.NewContainer(AddMyDependency())
	assert := assert.New(t)

	for range 3 {
		myDep, err := dependency.ResolveWithResolver[*MyDependency](c)

		assert.NoError(err)

		myDep.Count++

		t.Logf("%#v", myDep)
	}
}

func TestResolveWithResolverMissingDependency_Error(t *testing.T) {
	assert := assert.New(t)
	c := dependency.NewContainer()
	_, err := dependency.ResolveWithResolver[*MyDependency](c)

	assert.ErrorIs(err, dependency.ErrFactoryNotFound)
}

func TestResolveWithResolverWrongType_Error(t *testing.T) {
	assert := assert.New(t)
	c := dependency.NewContainer(AddMyDependencyWrongType())
	_, err := dependency.ResolveWithResolver[*MyDependency](c)

	assert.ErrorIs(err, dependency.ErrTypeConversion)
}
