package medium109

import (
	structures "leetcode-solutions/data-structures"
)

type list = structures.LinkedListNode
type node = structures.BinaryTreeNode

// time complexity: O(nlogn), space complexity: O(logn)，即结果平衡二叉树高度
func sortedListToBST(head *list) *node {
	if head == nil {
		return nil
	}

	// find middle node in the list
	slow, fast := head, head
	var prev *list
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &node{
		Val: slow.Val,
	}

	// recursion base case
	if head == slow {
		return root
	}

	if prev != nil {
		prev.Next = nil
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}
