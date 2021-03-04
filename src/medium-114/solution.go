package medium114

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func flatten0(root *node) {
	if root == nil {
		return
	}

	stack := []*node{root}
	var prev *node
	for len(stack) > 0 {
		index := len(stack) - 1
		current := stack[index]
		stack = stack[:index]

		// 相较于普通迭代法前序遍历多的步骤
		if prev != nil {
			prev.Left, prev.Right = nil, current
		}
		prev = current

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
}

func flatten1(root *node) {
	if root == nil {
		return
	}

	order := []*node{}
	stack := []*node{root}
	for len(stack) > 0 {
		index := len(stack) - 1
		current := stack[index]
		stack = stack[:index]
		order = append(order, current)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	for i := 1; i < len(order); i++ {
		prev := order[i-1]
		prev.Left = nil
		prev.Right = order[i]
	}
}
