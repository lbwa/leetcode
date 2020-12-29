package easy21

import structures "leetcode-solutions/data-structures"

type list = structures.LinkedListNode

func mergeTwoList(l1, l2 *list) *list {
	current1, current2 := l1, l2
	sentinel := &list{}
	current := sentinel

	for current1 != nil && current2 != nil {
		if current1.Val > current2.Val {
			current.Next = current2
			current2 = current2.Next
		} else {
			current.Next = current1
			current1 = current1.Next
		}
		current = current.Next
	}

	// 将未完成合并的链表部分直接合并
	if current1 == nil {
		current.Next = current2
	} else {
		current.Next = current1
	}

	return sentinel.Next
}
