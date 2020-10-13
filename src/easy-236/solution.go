package easy236

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func lowestCommonAncestor(root, p, q *node) (ancestor *node) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && p.Val > ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
