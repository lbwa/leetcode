package medium662

import (
	structures "leetcode-solutions/data-structures"
	"testing"
)

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	// [1,3,2,5,3,null,9]
	expect(t, widthOfBinaryTree(
		&structures.BinaryTreeNode{
			Val: 1,
			Left: &structures.BinaryTreeNode{
				Val: 3,
				Left: &structures.BinaryTreeNode{
					Val: 5,
				},
				Right: &structures.BinaryTreeNode{
					Val: 3,
				},
			},
			Right: &structures.BinaryTreeNode{
				Val: 2,
				Right: &structures.BinaryTreeNode{
					Val: 9,
				},
			},
		},
	), 4)

	// [1,3,null,5,3]
	expect(t, widthOfBinaryTree(
		&structures.BinaryTreeNode{
			Val: 1,
			Left: &structures.BinaryTreeNode{
				Val: 3,
				Left: &structures.BinaryTreeNode{
					Val: 5,
				},
				Right: &structures.BinaryTreeNode{
					Val: 3,
				},
			},
		},
	), 2)
}
