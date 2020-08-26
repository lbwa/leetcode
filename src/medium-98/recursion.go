package medium98

import "math"

func isValidBSTByRecursion(root *node) bool {
	return traversal(root, math.MinInt32, math.MaxInt32)
}

func traversal(root *node, low, high int) bool {
	if root == nil {
		return true
	}
	if (root.Val.(int) <= low) || (root.Val.(int) >= high) {
		return false
	}
	return traversal(root.Left, low, root.Val.(int)) && traversal(root.Right, root.Val.(int), high)
}
