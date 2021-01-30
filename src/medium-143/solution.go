package medium143

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

func reorderList(head *node) {
	if head == nil {
		return
	}

	nodes := []*node{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
