package medium662

import (
	structures "leetcode-solutions/data-structures"
)

type node struct {
	val   *structures.BinaryTreeNode
	index int
}

func widthOfBinaryTree(root *structures.BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	queue, answer := []node{node{val: root, index: 0}}, 0

	for len(queue) > 0 {
		answer = max(1+queue[len(queue)-1].index-queue[0].index, answer)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := shift(&queue)
			currentNode, index := current.val, current.index

			if currentNode.Left != nil {
				// 借鉴在 BST 及其相似结构通过顺序存储时，index 节点的左子节点为 2*index,  右子节点为 2*index + 1
				queue = append(queue, node{val: currentNode.Left, index: 2 * index})
			}
			if currentNode.Right != nil {
				queue = append(queue, node{val: currentNode.Right, index: 2*index + 1})
			}
		}
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func shift(nodes *[]node) node {
	deleted := (*nodes)[0]
	*nodes = (*nodes)[1:]
	return deleted
}
