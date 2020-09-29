package medium701

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func insertIntoBST(root *node, val int) *node {
	if root == nil {
		return &node{Val: val}
	}
	if root.Val.(int) > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val.(int) < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
