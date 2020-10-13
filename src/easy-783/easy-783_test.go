package easy783

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, minDiffInBST(&node{
		Val:   4,
		Left:  &node{Val: 2, Left: &node{Val: 1}, Right: &node{Val: 3}},
		Right: &node{Val: 6},
	}), 1)
}
