package medium24

import (
	structures "leetcode-solutions/data-structures"
)

func swapPairs(head *structures.LinkedListNode) *structures.LinkedListNode {
	if head == nil {
		return head
	}

	sentinel := &structures.LinkedListNode{
		Next: head,
	}
	current, prev := sentinel.Next, sentinel

	for current != nil {
		next := current.Next

		if next != nil {
			// 从链表中删除待反转节点
			current.Next = next.Next
			// 待反转节点指向被反转节点
			next.Next = current
			// 待反转节点重新接入到链表中
			prev.Next = next
		}

		prev = current
		current = current.Next
	}

	return sentinel.Next
}
