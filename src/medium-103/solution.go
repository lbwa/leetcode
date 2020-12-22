package medium103

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func zigzagLevelOrder(root *node) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		levelIndex := len(ans)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if levelIndex%2 > 0 {
				level[levelSize-1-i] = current.Val
			} else {
				level[i] = current.Val
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		ans = append(ans, level)
	}
	return
}
