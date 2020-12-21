package easy203

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func removeElement(head *node, val int) *node {
	sentinel := &node{Val: -1, Next: head}
	prev := sentinel
	current := sentinel.Next

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return sentinel.Next
}
