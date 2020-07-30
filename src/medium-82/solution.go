package medium82

import (
	structures "leetcode-solutions/data-structures"
)

func deleteDuplicated(head *structures.LinkedListNode) *structures.LinkedListNode {
	if head == nil {
		return head
	}

	sentinel := &structures.LinkedListNode{
		Val:  -1,
		Next: head,
	}
	current := sentinel

	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			// 直到下一项不等于下一项的下一项时
			now := current.Next
			for now.Next != nil && now.Val == now.Next.Val {
				now = now.Next
			}
			current.Next = now.Next
		} else {
			current = current.Next
		}
	}

	return sentinel.Next
}
