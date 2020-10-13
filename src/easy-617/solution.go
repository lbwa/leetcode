package easy617

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func mergeTrees(t1 *node, t2 *node) *node {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	treeNode := &node{Val: t1.Val + t2.Val}
	treeNode.Left = mergeTrees(t1.Left, t2.Left)
	treeNode.Right = mergeTrees(t1.Right, t2.Right)
	return treeNode
}
