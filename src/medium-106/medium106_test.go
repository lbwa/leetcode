package medium106

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}), &node{
		Val:  3,
		Left: &node{Val: 9},
		Right: &node{
			Val:   20,
			Left:  &node{Val: 15},
			Right: &node{Val: 7},
		},
	})
}
