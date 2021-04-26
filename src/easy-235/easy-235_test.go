package easy235

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	// [3,5,1,6,2,0,8,null,null,7,4]
	shared.Expect(t, lowestCommonAncestor(&node{
		Val: 3,
		Left: &node{
			Val: 5,
			Left: &node{
				Val: 6,
			},
			Right: &node{
				Val:   2,
				Left:  &node{Val: 7},
				Right: &node{Val: 4},
			},
		},
		Right: &node{
			Val: 1,
			Left: &node{
				Val: 0,
			},
			Right: &node{
				Val: 8,
			},
		},
	}, &node{Val: 5}, &node{Val: 1}), &node{
		Val: 3,
		Left: &node{
			Val: 5,
			Left: &node{
				Val: 6,
			},
			Right: &node{
				Val:   2,
				Left:  &node{Val: 7},
				Right: &node{Val: 4},
			},
		},
		Right: &node{
			Val: 1,
			Left: &node{
				Val: 0,
			},
			Right: &node{
				Val: 8,
			},
		},
	})
}
