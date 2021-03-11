package medium105

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}), &node{
		Val:  3,
		Left: &node{Val: 9},
		Right: &node{
			Val:   20,
			Left:  &node{Val: 15},
			Right: &node{Val: 7},
		},
	})
}
