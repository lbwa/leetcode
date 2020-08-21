package easy226

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, invertTree(&node{
		Val: 4,
		Left: &node{
			Val:   2,
			Left:  &node{Val: 1},
			Right: &node{Val: 3},
		},
		Right: &node{
			Val:   7,
			Left:  &node{Val: 6},
			Right: &node{Val: 9},
		},
	}), &node{
		Val: 4,
		Left: &node{
			Val:   7,
			Left:  &node{Val: 9},
			Right: &node{Val: 6},
		},
		Right: &node{
			Val:   2,
			Left:  &node{Val: 3},
			Right: &node{Val: 1},
		},
	})
}
