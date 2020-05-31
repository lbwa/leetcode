import { BinaryTreeNode } from 'data-structures/binary-tree'

export function maxDepth<V>(node: BinaryTreeNode<V> | null): number {
  if (node === null) return 0
  return Math.max(maxDepth(node.left), maxDepth(node.right)) + 1
}
