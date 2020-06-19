import { BinaryTreeNode } from 'data-structures/binary-tree'

/**
 * 核心是 https://algs4.cs.princeton.edu/32bst/ 算法4 中提及的 BST 的 put 节点的操作
 */
export function bstFromPreOrder(
  preOrder: number[]
): BinaryTreeNode<number> | null {
  const head = new BinaryTreeNode(preOrder[0])

  for (let i = 1; i < preOrder.length; i++) {
    putBSTFromNode(head, preOrder[i])
  }

  return head
}

function putBSTFromNode<V extends number>(
  node: BinaryTreeNode<V> | null,
  value: V
) {
  if (node === null) return new BinaryTreeNode(value)

  const comparison = value - node.value
  if (comparison > 0) {
    node.right = putBSTFromNode(node.right, value)
  } else if (comparison < 0) {
    node.left = putBSTFromNode(node.left, value)
  } else {
    node.value = value
  }

  return node
}
