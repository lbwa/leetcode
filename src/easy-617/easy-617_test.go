package easy617

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, mergeTrees(&node{
		Val: 1,
		Left: &node{
			Val:  3,
			Left: &node{Val: 5},
		},
		Right: &node{Val: 2},
	}, &node{
		Val: 2,
		Left: &node{
			Val:   1,
			Right: &node{Val: 4},
		},
		Right: &node{
			Val:   3,
			Right: &node{Val: 7},
		},
	}), &node{
		Val: 3,
		Left: &node{
			Val:   4,
			Left:  &node{Val: 5},
			Right: &node{Val: 4},
		},
		Right: &node{
			Val: 5,
			Right: &node{
				Val: 7,
			},
		},
	})
}
