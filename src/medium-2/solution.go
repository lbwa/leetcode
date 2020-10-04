package medium2

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func addTwoNumbers(l1 *node, l2 *node) *node {
	sentinel := &node{}
	current, current1, current2, carry := sentinel, l1, l2, 0
	for current1 != nil || current2 != nil {
		var sum int
		if current1 != nil {
			sum += current1.Val
			current1 = current1.Next
		}
		if current2 != nil {
			sum += current2.Val
			current2 = current2.Next
		}
		sum += carry
		carry = sum / 10
		current.Next = &node{Val: sum % 10}
		current = current.Next
	}
	if carry != 0 {
		current.Next = &node{Val: carry}
	}
	return sentinel.Next
}
