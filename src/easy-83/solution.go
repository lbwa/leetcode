package easy83

import (
	structures "leetcode-solutions/data-structures"
)

func deleteDuplicated(head *structures.LinkedListNode) *structures.LinkedListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head.Next

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

	slow.Next = nil

	return head
}
