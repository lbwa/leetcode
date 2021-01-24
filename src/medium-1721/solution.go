package medium1721

import structures "leetcode-solutions/data-structures"

type node = structures.LinkedListNode

// 交换两节点
func swapNodes(head *node, k int) *node {
	if head == nil {
		return head
	}

	sentinel := &node{Next: head}
	current := sentinel

	for k > 1 {
		current = current.Next
		k--
	}

	left, right := current, sentinel

	for current.Next.Next != nil {
		current = current.Next
		right = right.Next
	}

	// 待交换的两节点
	leftNext := right.Next
	rightNext := left.Next

	// 交换两节点
	left.Next = leftNext
	right.Next = rightNext

	// 交换的两节点的后续节点
	leftNextNext := rightNext.Next
	rightNextNext := leftNext.Next

	// 恢复交换的两节点的后续节点
	leftNext.Next = leftNextNext
	rightNext.Next = rightNextNext

	return sentinel.Next
}

// 仅交换两节点的值
func swapNodesVal(head *node, k int) *node {
	if head == nil {
		return head
	}

	current := head

	for k > 1 {
		current = current.Next
		k--
	}

	fast, slow := head, current

	for current.Next != nil {
		current = current.Next
		fast = fast.Next
	}

	fast.Val, slow.Val = slow.Val, fast.Val

	return head
}
