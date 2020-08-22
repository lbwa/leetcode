package medium105

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func buildTree(preorder, inorder []int) *node {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	val := preorder[0]
	var i int
	for i = range inorder {
		if inorder[i] == val {
			break
		}
	}
	return &node{
		Val: val,
		// 保证传入的 preorder 和 inorder 长度相等，才能保证每次都能找到 root 值
		// 故 preorder 结束项为 i+1
		Left:  buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}
