package medium189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	a0 := []int{1, 2, 3, 4, 5, 6, 7}
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	b0 := []int{-1, -100, 3, 99}
	b1 := []int{-1, -100, 3, 99}

	rotate(a0, 3)
	rotate(b0, 2)
	rotate0(a1, 3)
	rotate0(b1, 2)
	assert.Equal([]int{5, 6, 7, 1, 2, 3, 4}, a0)
	assert.Equal([]int{3, 99, -1, -100}, b0)
	assert.Equal([]int{5, 6, 7, 1, 2, 3, 4}, a1)
	assert.Equal([]int{3, 99, -1, -100}, b1)
}
