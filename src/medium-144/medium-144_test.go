package medium144

import (
	structures "leetcode-solutions/data-structures"
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, preorderTraversal(&structures.BinaryTreeNode{
		Val: 1,
		Right: &structures.BinaryTreeNode{
			Val: 2,
			Left: &structures.BinaryTreeNode{
				Val: 3,
			},
		},
	}), []int{1, 2, 3})
}
