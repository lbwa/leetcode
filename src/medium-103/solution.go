package medium103

import (
	structures "leetcode-solutions/data-structures"
)

func zigzagLevelOrder(root *structures.BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue, answer := []*structures.BinaryTreeNode{root}, [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		answer = append(answer, []int{})
		levelIndex := len(answer) - 1
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			answer[levelIndex] = append(answer[levelIndex], current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		if levelIndex%2 != 0 {
			currentLevel := answer[levelIndex]
			for i, j := 0, len(currentLevel)-1; i < j; i, j = i+1, j-1 {
				currentLevel[i], currentLevel[j] = currentLevel[j], currentLevel[i]
			}
		}
	}

	return answer
}
