package medium19

import (
	structures "leetcode-solutions/data-structures"
)

func removeNthFromEnd(head *structures.LinkedListNode, n int) *structures.LinkedListNode {
	sentinel := &structures.LinkedListNode{
		Next: head,
	}
	slow, fast := sentinel, sentinel

	for i := 1; i <= n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return sentinel.Next
}
