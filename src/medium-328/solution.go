package medium328

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func oddEvenList(head *node) *node {
	if head == nil {
		return head
	}

	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
