package easy160

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func getIntersectionNode(headA, headB *node) *node {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
