package medium114

import (
	structures "leetcode-solutions/data-structures"
)

func flatten0(root *structures.BinaryTreeNode) {
	if root == nil {
		return
	}

	stack := []*structures.BinaryTreeNode{root}
	var prev *structures.BinaryTreeNode

	for len(stack) > 0 {
		currentIndex := len(stack) - 1
		current := stack[currentIndex]
		stack = stack[:currentIndex]

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

func flatten1(root *structures.BinaryTreeNode) {
	if root == nil {
		return
	}

	stack := []*structures.BinaryTreeNode{root}
	preorderNodes := []*structures.BinaryTreeNode{}
	for len(stack) > 0 {
		currentIndex := len(stack) - 1
		current := stack[currentIndex]
		stack, preorderNodes = stack[:currentIndex], append(preorderNodes, current)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	for i := 1; i < len(preorderNodes); i++ {
		prev, node := preorderNodes[i-1], preorderNodes[i]
		prev.Left, prev.Right = nil, node
	}
}
