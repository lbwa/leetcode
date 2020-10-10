package medium142

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func detectCycle(head *node) *node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	// 无环
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 有环
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
