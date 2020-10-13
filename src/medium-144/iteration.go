package medium144

import (
	structures "leetcode-solutions/data-structures"
)

func preorderTraversal(root *structures.BinaryTreeNode) []int {
	answer := []int{}
	if root == nil {
		return answer
	}

	stack := []*structures.BinaryTreeNode{root}
	for len(stack) > 0 {
		currentIndex := len(stack) - 1
		current := stack[currentIndex]
		stack, answer = stack[:currentIndex], append(answer, current.Val)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return answer
}
