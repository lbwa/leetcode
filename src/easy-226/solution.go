package easy226

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func invertTree(root *node) *node {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
