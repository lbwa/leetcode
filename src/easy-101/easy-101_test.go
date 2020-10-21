package easy101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, isSymmetric(&node{
		Val:   1,
		Left:  &node{Val: 2, Left: &node{Val: 3}, Right: &node{Val: 4}},
		Right: &node{Val: 2, Left: &node{Val: 4}, Right: &node{Val: 3}},
	}), true)

	assert.Equal(t, isSymmetric(&node{
		Val:   1,
		Left:  &node{Val: 2, Right: &node{Val: 3}},
		Right: &node{Val: 2, Right: &node{Val: 3}},
	}), false)
}
