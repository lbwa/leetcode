package solution

import (
	structures "leetcode-solutions/data-structures"
)

// IterativeInOrderTraversal is medium 94 solution
func IterativeInOrderTraversal(root *structures.BinaryTreeNode) []int {
	stack := []*structures.BinaryTreeNode{}
	answer := []int{}

	if root == nil {
		return answer
	}

	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		latest := len(stack) - 1
		current = stack[latest]
		stack = stack[:latest]
		// type assertion
		// https://tour.golang.org/methods/15
		answer = append(answer, current.Val.(int))
		current = current.Right
	}

	return answer
}
