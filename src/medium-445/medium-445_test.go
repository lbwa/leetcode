package medium445

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, addTwoNumbers(&node{
		Val: 7,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 4,
				Next: &node{
					Val: 3,
				},
			},
		},
	}, &node{
		Val: 5,
		Next: &node{
			Val: 6,
			Next: &node{
				Val: 4,
			},
		},
	}), &node{
		Val: 7,
		Next: &node{
			Val: 8,
			Next: &node{
				Val: 0,
				Next: &node{
					Val: 7,
				},
			},
		},
	})
}
