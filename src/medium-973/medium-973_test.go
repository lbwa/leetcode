package medium973

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{[]int{-2, 2}}, kClosest([][]int{[]int{1, 3}, []int{-2, 2}}, 1))
	assert.Equal([][]int{[]int{3, 3}, []int{-2, 4}}, kClosest([][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2))
	assert.Equal([][]int{[]int{-2, 2}}, kClosestPQ([][]int{[]int{1, 3}, []int{-2, 2}}, 1))
	assert.Equal([][]int{[]int{-2, 4}, []int{3, 3}}, kClosestPQ([][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2))
}
