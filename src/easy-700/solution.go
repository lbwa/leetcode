package easy700

import structures "leetcode-solutions/data-structures"

type node = structures.BinaryTreeNode

// 因为是 BST 故借鉴其左子树各个节点小于根节点，右子树反之的特性，而使用二分法得到目标子树
func searchBST(root *node, val int) *node {
	if root == nil || (root.Val == val) {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
