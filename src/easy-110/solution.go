package easy110

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

// based on top-to-bottom dfs
// func isBalanced(root *node) bool {
// 	if root == nil {
// 		return true
// 	}

// 	return abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
// }

// func height(root *node) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return max(height(root.Left), height(root.Right)) + 1
// }

// based on bottom-to-top dfs
func isBalanced(root *node) bool {
	return height(root) > 0
}

func height(root *node) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	if left == -1 {
		return -1
	}
	right := height(root.Right)
	if right == -1 || abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
