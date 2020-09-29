package medium701

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, insertIntoBST(&node{
		Val: 4,
		Left: &node{
			Val:   2,
			Left:  &node{Val: 1},
			Right: &node{Val: 3},
		},
		Right: &node{Val: 7},
	}, 5), &node{
		Val: 4,
		Left: &node{
			Val:   2,
			Left:  &node{Val: 1},
			Right: &node{Val: 3},
		},
		Right: &node{Val: 7, Left: &node{Val: 5}},
	})
}
