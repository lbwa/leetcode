package medium48

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	matrix0 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix0)
	assert.Equal([][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, matrix0)
}
