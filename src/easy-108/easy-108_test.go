package easy108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, sortedArrayToBST([]int{-10, -3, 0, 5, 9}), &node{
		Val:   0,
		Left:  &node{Val: -10, Right: &node{Val: -3}},
		Right: &node{Val: 5, Right: &node{Val: 9}},
	})
}
