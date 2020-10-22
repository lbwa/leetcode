package easy234

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func isPalindrome(head *node) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	left, right := head, reverseList(slow)

	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func reverseList(head *node) *node {
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
