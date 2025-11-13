package dependency_test

import (
	"testing"

	"github.com/securehaven/dependency"
	"github.com/stretchr/testify/assert"
)

func TestName_Success(t *testing.T) {
	expected := "*dependency_test.MyDependency"
	received := dependency.Name(new(MyDependency))

	assert.Equal(t, expected, received)
}
