package medium113

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{
		[]int{5, 4, 11, 2},
		[]int{5, 8, 4, 5},
	}, pathSum(&node{
		Val: 5,
		Left: &node{
			Val: 4,
			Left: &node{
				Val:   11,
				Left:  &node{Val: 7},
				Right: &node{Val: 2},
			},
		},
		Right: &node{
			Val:  8,
			Left: &node{Val: 13},
			Right: &node{
				Val:   4,
				Left:  &node{Val: 5},
				Right: &node{Val: 1},
			},
		},
	}, 22))
}
