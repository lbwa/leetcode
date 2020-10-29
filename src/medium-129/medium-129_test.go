package medium129

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(25, sumNumbers(&node{
		Val:   1,
		Left:  &node{Val: 2},
		Right: &node{Val: 3},
	}))
	assert.Equal(25, sumNumbers0(&node{
		Val:   1,
		Left:  &node{Val: 2},
		Right: &node{Val: 3},
	}))
}
