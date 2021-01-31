package medium86

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func partition(head *node, x int) *node {
	if head == nil {
		return head
	}

	lessHead, greatHead := &node{}, &node{}
	less, great := lessHead, greatHead

	current := head
	for current != nil {
		if current.Val < x {
			less.Next = current
			less = less.Next
		} else {
			great.Next = current
			great = great.Next
		}
		current = current.Next
	}
	less.Next = greatHead.Next
	great.Next = nil
	return lessHead.Next
}
