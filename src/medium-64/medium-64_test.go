package medium64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}), 7)
	assert.Equal(minPathSum([][]int{
		{5},
	}), 5)
	assert.Equal(minPathSum([][]int{}), 0)
	assert.Equal(minPathSum([][]int{{}}), 0)
}
