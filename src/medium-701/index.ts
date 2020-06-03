import { BinaryTreeNode } from 'data-structures/binary-tree'

export function insertIntoBST(
  node: BinaryTreeNode<number> | null,
  val: number
): BinaryTreeNode<number> {
  if (node === null) {
    return new BinaryTreeNode(val, null, null)
  }

  const compare = node.value - val
  if (compare > 0) {
    node.left = insertIntoBST(node.left, val)
  } else if (compare < 0) {
    node.right = insertIntoBST(node.right, val)
  } else {
    node.value = val
  }

  return node
}
