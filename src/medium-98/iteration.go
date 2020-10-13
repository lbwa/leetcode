package medium98

import (
	structures "leetcode-solutions/data-structures"
	"math"
)

type node = structures.BinaryTreeNode

func isValidBSTByIteration(root *node) bool {
	if root == nil {
		return true
	}

	stack, current, prevNodeVal := []*node{}, root, math.MinInt32
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		topFrameIndex := len(stack) - 1
		current = stack[topFrameIndex]
		stack = stack[:topFrameIndex]

		if current.Val <= prevNodeVal {
			return false
		}
		prevNodeVal = current.Val

		current = current.Right
	}
	return true
}
