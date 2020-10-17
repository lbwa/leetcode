package medium46

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, permute([]int{1, 2, 3}), [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}, "Should be equal")

	assert.Equal(t, permute([]int{5, 4, 6, 2}), [][]int{
		[]int{5, 4, 6, 2},
		[]int{5, 4, 2, 6},
		[]int{5, 6, 4, 2},
		[]int{5, 6, 2, 4},
		[]int{5, 2, 4, 6},
		[]int{5, 2, 6, 4},
		[]int{4, 5, 6, 2},
		[]int{4, 5, 2, 6},
		[]int{4, 6, 5, 2},
		[]int{4, 6, 2, 5},
		[]int{4, 2, 5, 6},
		[]int{4, 2, 6, 5},
		[]int{6, 5, 4, 2},
		[]int{6, 5, 2, 4},
		[]int{6, 4, 5, 2},
		[]int{6, 4, 2, 5},
		[]int{6, 2, 5, 4},
		[]int{6, 2, 4, 5},
		[]int{2, 5, 4, 6},
		[]int{2, 5, 6, 4},
		[]int{2, 4, 5, 6},
		[]int{2, 4, 6, 5},
		[]int{2, 6, 5, 4},
		[]int{2, 6, 4, 5},
	})
}
