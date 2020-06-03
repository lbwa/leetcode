import { BinaryTreeNode } from 'data-structures/binary-tree'

export function kthSmallest(
  node: BinaryTreeNode<number> | null,
  k: number
): number {
  // 实现了以 0 为始的 BST 选择函数，故减去 1
  return select(node, k - 1)
}

// 基于 `算法 4` 中的 BST select 方法实现，始于 0
function select(node: BinaryTreeNode<number> | null, k: number): number {
  if (node === null) return 0

  const leftSubtreeSize = getNodeSize(node.left)
  if (k < leftSubtreeSize) {
    return select(node.left, k)
  } else if (k > leftSubtreeSize) {
    return select(node.right, k - leftSubtreeSize - 1)
  } else {
    return node.value
  }
}

function getNodeSize(node: BinaryTreeNode<number> | null): number {
  if (node === null) return 0

  return 1 + getNodeSize(node.left) + getNodeSize(node.right)
}
