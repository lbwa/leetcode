package easy404

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, sumOfLeftLeaves(&node{
		Val:   1,
		Left:  &node{Val: 2, Left: &node{Val: 4}, Right: &node{Val: 5}},
		Right: &node{Val: 3},
	}), 4)
	shared.Expect(t, sumOfLeftLeaves(&node{
		Val:   3,
		Left:  &node{Val: 9},
		Right: &node{Val: 20, Left: &node{Val: 15}, Right: &node{Val: 7}},
	}), 24)
}
