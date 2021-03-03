package easy1137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, tribonacci(4))
	assert.Equal(1389537, tribonacci(25))

	assert.Equal(4, tribonacci0(4))
	assert.Equal(1389537, tribonacci0(25))
}
