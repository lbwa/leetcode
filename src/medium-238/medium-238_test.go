package medium238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{24, 12, 8, 6}, productExceptSelfOn([]int{1, 2, 3, 4}))
	assert.Equal([]int{24, 12, 8, 6}, productExceptSelfO1([]int{1, 2, 3, 4}))
}
