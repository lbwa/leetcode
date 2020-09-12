package easy637

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func averageOfLevels(root *node) (answer []float64) {
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			levelSum += current.Val.(int)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		answer = append(answer, float64(levelSum)/float64(levelSize))
	}
	return
}
