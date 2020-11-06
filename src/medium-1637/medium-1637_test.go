package medium1637

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, maxWidthOfVerticalArea([][]int{
		[]int{8, 7},
		[]int{9, 9},
		[]int{7, 4},
		[]int{9, 7},
	}))
	assert.Equal(3, maxWidthOfVerticalArea([][]int{
		[]int{3, 1},
		[]int{9, 0},
		[]int{1, 0},
		[]int{1, 4},
		[]int{5, 3},
		[]int{8, 8},
	}))
}
