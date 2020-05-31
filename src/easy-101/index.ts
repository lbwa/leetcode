import { BinaryTreeNode } from 'data-structures/binary-tree'
import { isDef } from 'src/shared'

export function isSymmetric<V>(node: BinaryTreeNode<V> | null) {
  return compareNodes(node, node)
}

function compareNodes<V>(
  a: BinaryTreeNode<V> | null,
  b: BinaryTreeNode<V> | null
): boolean {
  if (!isDef(a) && !isDef(b)) return true
  if (!isDef(a) || !isDef(b)) return false
  return (
    a.value === b.value &&
    compareNodes(a.left, b.right) &&
    compareNodes(a.right, b.left)
  )
}
