package easy141

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func hasCycle(head *node) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
