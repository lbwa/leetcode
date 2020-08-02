package medium114

import (
	structures "leetcode-solutions/data-structures"
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	flatten0(nil)
	flatten1(nil)
	tree0 := &structures.BinaryTreeNode{
		Val: 1,
		Left: &structures.BinaryTreeNode{
			Val: 2,
			Left: &structures.BinaryTreeNode{
				Val: 3,
			},
			Right: &structures.BinaryTreeNode{
				Val: 4,
			},
		},
		Right: &structures.BinaryTreeNode{
			Val: 5,
			Right: &structures.BinaryTreeNode{
				Val: 6,
			},
		},
	}
	tree1 := &structures.BinaryTreeNode{
		Val: 1,
		Left: &structures.BinaryTreeNode{
			Val: 2,
			Left: &structures.BinaryTreeNode{
				Val: 3,
			},
			Right: &structures.BinaryTreeNode{
				Val: 4,
			},
		},
		Right: &structures.BinaryTreeNode{
			Val: 5,
			Right: &structures.BinaryTreeNode{
				Val: 6,
			},
		},
	}
	flatten0(tree0)
	flatten1(tree1)
	shared.Expect(t, tree0, &structures.BinaryTreeNode{
		Val: 1,
		Right: &structures.BinaryTreeNode{
			Val: 2,
			Right: &structures.BinaryTreeNode{
				Val: 3,
				Right: &structures.BinaryTreeNode{
					Val: 4,
					Right: &structures.BinaryTreeNode{
						Val: 5,
						Right: &structures.BinaryTreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	})
	shared.Expect(t, tree1, &structures.BinaryTreeNode{
		Val: 1,
		Right: &structures.BinaryTreeNode{
			Val: 2,
			Right: &structures.BinaryTreeNode{
				Val: 3,
				Right: &structures.BinaryTreeNode{
					Val: 4,
					Right: &structures.BinaryTreeNode{
						Val: 5,
						Right: &structures.BinaryTreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	})
}
