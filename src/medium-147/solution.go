package medium147

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func insertionSortlist(head *node) *node {
	if head == nil {
		return head
	}
	sentinel := &node{Next: head}
	sortedPartTail, current := head, head.Next
	for current != nil {
		if sortedPartTail.Val <= current.Val {
			sortedPartTail = sortedPartTail.Next
		} else {
			prev := sentinel
			for prev.Next.Val <= current.Val {
				prev = prev.Next
			}
			sortedPartTail.Next = current.Next // remove current from list

			// resume link
			current.Next = prev.Next
			prev.Next = current
		}
		current = sortedPartTail.Next
	}
	return sentinel.Next
}
