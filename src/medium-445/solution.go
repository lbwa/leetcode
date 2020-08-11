package medium445

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func addTwoNumbers(l1 *node, l2 *node) *node {
	sentinel := &node{}
	current, carry := l1, 0
	stack1, stack2 := []*node{}, []*node{}
	for current != nil {
		stack1 = append(stack1, current)
		current = current.Next
	}
	current = l2
	for current != nil {
		stack2 = append(stack2, current)
		current = current.Next
	}

	for len(stack1) > 0 || len(stack2) > 0 {
		a, b := 0, 0
		if len(stack1) > 0 {
			last := len(stack1) - 1
			a = stack1[last].Val
			stack1 = stack1[:last]
		}
		if len(stack2) > 0 {
			last := len(stack2) - 1
			b = stack2[last].Val
			stack2 = stack2[:last]
		}
		sum := a + b + carry
		carry = sum / 10
		sentinel.Next = &node{Val: sum % 10, Next: sentinel.Next}
	}

	if carry > 0 {
		sentinel.Next = &node{Val: carry, Next: sentinel.Next}
	}

	return sentinel.Next
}
