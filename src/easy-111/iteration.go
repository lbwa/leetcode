package easy111

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func iterativeMinDepth(root *node) int {
	if root == nil {
		return 0
	}

	queue, depth := []*node{root}, 0
	for len(queue) > 0 {
		levelSize := len(queue)
		depth++
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left == nil && current.Right == nil {
				return depth
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return depth
}
