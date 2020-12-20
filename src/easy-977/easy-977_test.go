package easy977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{0, 1, 9, 16, 100}, sortedSqares([]int{-4, -1, 0, 3, 10}))
	assert.Equal([]int{4, 9, 9, 49, 121}, sortedSqares([]int{-7, -3, 2, 3, 11}))
}
