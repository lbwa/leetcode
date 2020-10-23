package easy112

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func hasPathSum(root *node, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum-root.Val == 0
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
