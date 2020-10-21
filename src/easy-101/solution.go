package easy101

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func isSymmetric(root *node) bool {
	var traversal func(a, b *node) bool
	traversal = func(a, b *node) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.Val == b.Val && traversal(a.Left, b.Right) && traversal(a.Right, b.Left)
	}
	return traversal(root, root)
}
