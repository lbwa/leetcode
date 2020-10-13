package medium98

import "math"

func isValidBSTByRecursion(root *node) bool {
	return traversal(root, math.MinInt32, math.MaxInt32)
}

func traversal(root *node, low, high int) bool {
	if root == nil {
		return true
	}
	if (root.Val <= low) || (root.Val >= high) {
		return false
	}
	return traversal(root.Left, low, root.Val) && traversal(root.Right, root.Val, high)
}
