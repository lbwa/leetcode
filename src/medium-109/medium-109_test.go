package medium109

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, sortedListToBST(&list{
		Val: -10,
		Next: &list{
			Val: -3,
			Next: &list{
				Val: 0,
				Next: &list{
					Val: 5,
					Next: &list{
						Val: 9,
					},
				},
			},
		},
	}), &node{
		Val: 0,
		Left: &node{
			Val: -3,
			Left: &node{
				Val: -10,
			},
		},
		Right: &node{
			Val: 9,
			Left: &node{
				Val: 5,
			},
		},
	})
}
