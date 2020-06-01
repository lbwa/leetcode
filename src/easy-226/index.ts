import { BinaryTreeNode } from 'data-structures/binary-tree'

export function invertTree<V>(root: BinaryTreeNode<V> | null) {
  if (root === null) return null

  const temp = root.left
  root.left = root.right
  root.right = temp

  invertTree(root.left)
  invertTree(root.right)

  return root
}
