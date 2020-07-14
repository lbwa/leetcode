package easy203

import (
	structures "leetcode-solutions/data-structures"
)

// RemoveElement is used to delete items from singly linked list
func RemoveElement(head *structures.LinkedListNode, val int) *structures.LinkedListNode {
	sentinel := &structures.LinkedListNode{Val: -1, Next: head}
	prev := sentinel
	current := sentinel.Next

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return sentinel.Next
}
