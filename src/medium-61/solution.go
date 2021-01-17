package medium61

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func rotateRight(head *node, k int) *node {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	current := head
	for current.Next != nil {
		length++
		current = current.Next
	}

	k = k % length
	current.Next = head
	current = head
	for i := 0; i < length-k-1; i++ {
		current = current.Next
	}

	head = current.Next
	current.Next = nil
	return head
}
