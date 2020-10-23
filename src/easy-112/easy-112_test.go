package easy112

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(hasPathSum(&node{
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
				Val:  4,
				Left: &node{Val: 1},
			},
		},
	}, 22), true)
}
