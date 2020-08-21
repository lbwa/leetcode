package easy111

import "math"

func recursiveMinDepth(root *node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt32
	if root.Left != nil {
		depth = min(recursiveMinDepth(root.Left), depth)
	}
	if root.Right != nil {
		depth = min(recursiveMinDepth(root.Right), depth)
	}
	return depth + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
