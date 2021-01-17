package medium92

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func reverseBetween(head *node, m, n int) *node {
	if head == nil {
		return head
	}
	sentinel := &node{Next: head}

	prev := sentinel
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	current := prev.Next
	// 头插法反转链表部分区间
	for i := m; i < n; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return sentinel.Next
}
