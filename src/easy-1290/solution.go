package easy1290

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func getDecimalValue(head *node) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}

	return sum
}

func getDecimalValue0(head *node) int {
	sum := 0
	for head != nil {
		sum = (sum << 1) + head.Val
		head = head.Next
	}

	return sum
}
