package medium1038

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func bstToGst(root *node) *node {
	var prev int
	var dfs func(root *node) *node
	dfs = func(root *node) *node {
		if root == nil {
			return root
		}
		dfs(root.Right)
		prev += root.Val.(int)
		root.Val = prev
		dfs(root.Left)
		return root
	}
	return root
}
