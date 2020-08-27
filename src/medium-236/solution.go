package medium236

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func lowestCommonAncestor(root, p, q *node) *node {
	if root == nil || root == p || root == q {
		return root
	}

	// 递归遍历 left 和 right 子树中，是否包含 p **或** q 节点
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右子树均不包含 p 或 q，那么 p 和 q 并不存在于树种
	if left == nil && right == nil {
		return nil
	}

	// 若左右子树均返回节点，那么最近公共祖先节点即为当前 root 节点
	// 以上基于后序遍历，那么始终从最底层的叶子节点开始更新，那么在所有满足条件的祖先
	// 节点种一定是最近的组件节点首先被访问到
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}
