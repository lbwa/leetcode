package medium1669

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func mergeInBetween(list1 *node, a, b int, list2 *node) *node {
	if list1 == nil {
		return list1
	}

	sentinel := &node{Next: list1}
	start, end := list1, list1
	for b > -1 {
		if a > 1 {
			start = start.Next
			a--
		}
		end = end.Next
		b--
	}

	start.Next = list2
	if list2 != nil {
		current := list2
		for current.Next != nil {
			current = current.Next
		}
		current.Next = end
	}

	return sentinel.Next
}
