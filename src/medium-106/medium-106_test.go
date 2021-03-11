package medium106

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}), &node{
		Val:  3,
		Left: &node{Val: 9},
		Right: &node{
			Val:   20,
			Left:  &node{Val: 15},
			Right: &node{Val: 7},
		},
	})
}
