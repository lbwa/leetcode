package easy206

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func reverseListIteration(head *node) *node {
	if head == nil {
		return head
	}

	sentinel := &node{Next: head}

	for head.Next != nil {
		next := head.Next
		head.Next = head.Next.Next
		next.Next = sentinel.Next
		sentinel.Next = next
	}

	return sentinel.Next
}

func reverseListRecursion(head *node) *node {
	if head == nil || head.Next == nil {
		return head
	}

	tail := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}
