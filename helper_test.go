package dependency_test

import (
	"testing"

	"github.com/securehaven/dependency"
	"github.com/stretchr/testify/assert"
)

func TestHelperName_Success(t *testing.T) {
	expected := "*dependency_test.MyDependency"
	received := dependency.Name(new(MyDependency))

	assert.Equal(t, expected, received)
}

func TestHelperStart_Success(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependency())
	myDepResolved := dependency.Start[*MyDependency](container)

	assert.NoError(myDepResolved.Err)
	assert.IsType(new(MyDependency), myDepResolved.Value)
}

func TestHelperStart_Error(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependencyError())
	myDepResolved := dependency.Start[*MyDependency](container)

	assert.ErrorIs(myDepResolved.Err, ErrMyDependency)
}

func TestHelperThen_Success(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependency())
	firstDepResolved := dependency.Start[*MyDependency](container)
	secondDepResolved := dependency.Then[*MyDependency](container, firstDepResolved)

	assert.NoError(firstDepResolved.Err)
	assert.NoError(secondDepResolved.Err)
	assert.IsType(new(MyDependency), firstDepResolved.Value)
	assert.IsType(new(MyDependency), secondDepResolved.Value)
}

func TestHelperThen_Error(t *testing.T) {
	assert := assert.New(t)
	container := dependency.NewContainer(AddMyDependencyError())
	firstDepResolved := dependency.Start[*MyDependency](container)
	secondDepResolved := dependency.Then[*MyDependency](container, firstDepResolved)

	assert.ErrorIs(firstDepResolved.Err, ErrMyDependency)
	assert.ErrorIs(secondDepResolved.Err, ErrMyDependency)
}
