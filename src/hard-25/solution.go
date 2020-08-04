package hard25

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.LinkedListNode

func reverseKgroup(head *node, k int) *node {
	if head == nil {
		return head
	}
	sentinel := &node{
		Next: head,
	}
	current, prev := sentinel.Next, sentinel

	for current != nil {
		// 是否应该执行 k 区间反转
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 当不能执行 k 个节点内的区间反转时，表示剩下接待不足 k 个，那么返回头节点
				return sentinel.Next
			}
		}

		// 反转 k 个节点
		for i := 0; i < k-1; i++ {
			next := current.Next
			current.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		// 此时 current 为 k 个节点区间内的 tail 尾节点
		prev = current
		// 进入下一个 k 区间
		current = current.Next
	}

	return sentinel.Next
}
