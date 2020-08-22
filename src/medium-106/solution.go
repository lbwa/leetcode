package medium106

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func buildTree(inorder, postorder []int) *node {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	// based on unique value
	var i int
	for i = range inorder {
		if inorder[i] == rootVal {
			break
		}
	}
	// 找到区分左右子树的值，那么此时该值左右两边分别对应左右子树
	return &node{
		Val: rootVal,
		// 划分左边结果数组，构建左子树
		Left: buildTree(inorder[:i], postorder[:i]),
		// 反之，构建右子树
		Right: buildTree(inorder[i+1:], postorder[i:(len(postorder)-1)]),
	}
	// 每次递归后，inorder 中都会移除第 i 项(即 root 值)，postorder 中都会移除末项（即 root 值）
}
