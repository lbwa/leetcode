package easy236

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func lowestCommonAncestor(root, p, q *node) (ancestor *node) {
	ancestor = root
	for {
		if p.Val.(int) < ancestor.Val.(int) && p.Val.(int) > ancestor.Val.(int) {
			ancestor = ancestor.Left
		} else if p.Val.(int) > ancestor.Val.(int) && q.Val.(int) > ancestor.Val.(int) {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
