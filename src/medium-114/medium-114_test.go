package medium114

import (
	structures "leetcode-solutions/data-structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
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
	assert.Equal(t, tree0, &structures.BinaryTreeNode{
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
	assert.Equal(t, tree1, &structures.BinaryTreeNode{
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
