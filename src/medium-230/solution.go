package medium230

import structures "leetcode-solutions/data-structures"

type node = structures.BinaryTreeNode

// 基于二叉搜索树的中序遍历为 小到大 升序的特点
func kthSmallest(root *node, k int) int {
	stack := []*node{}
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		topFrame := len(stack) - 1
		current = stack[topFrame]
		stack = stack[:topFrame]
		k--
		if k == 0 {
			return current.Val
		}
		current = current.Right
	}

	return *(*int)(nil)
}
