package medium889

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}), &node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 4,
			},
			Right: &node{
				Val: 5,
			},
		},
		Right: &node{
			Val: 3,
			Left: &node{
				Val: 6,
			},
			Right: &node{
				Val: 7,
			},
		},
	})
}
