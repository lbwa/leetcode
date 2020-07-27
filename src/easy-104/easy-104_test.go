package easy104

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
	expect(t, maxDepth(&structures.BinaryTreeNode{
		Val:  3,
		Left: &structures.BinaryTreeNode{Val: 9},
		Right: &structures.BinaryTreeNode{
			Val:   20,
			Left:  &structures.BinaryTreeNode{Val: 15},
			Right: &structures.BinaryTreeNode{Val: 7},
		},
	}), 3)
	expect(t, maxDepthWithLevelOrder(&structures.BinaryTreeNode{
		Val:  3,
		Left: &structures.BinaryTreeNode{Val: 9},
		Right: &structures.BinaryTreeNode{
			Val:   20,
			Left:  &structures.BinaryTreeNode{Val: 15},
			Right: &structures.BinaryTreeNode{Val: 7},
		},
	}), 3)
}
