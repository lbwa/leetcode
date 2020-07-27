package easy104

import (
	structures "leetcode-solutions/data-structures"
)

func maxDepth(root *structures.BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepthWithLevelOrder(root *structures.BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	queue, depth := []*structures.BinaryTreeNode{root}, 0

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
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
