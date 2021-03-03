package easy509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, fib(2))
	assert.Equal(2, fib(3))
	assert.Equal(3, fib(4))

	assert.Equal(1, fib0(2))
	assert.Equal(2, fib0(3))
	assert.Equal(3, fib0(4))

}
