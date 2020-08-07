package easy100

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func isSameTree(p, q *node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
