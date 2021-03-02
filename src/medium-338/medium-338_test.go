package medium338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 1, 1}, countBits(2))
	assert.Equal([]int{0, 1, 1, 2, 1, 2}, countBits(5))
}
