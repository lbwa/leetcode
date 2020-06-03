import { BinaryTreeNode } from 'data-structures/binary-tree'

export function searchBST(
  node: BinaryTreeNode<number> | null,
  val: number
): BinaryTreeNode<number> | null {
  if (node === null) return null

  const compare = node.value - val
  // 因为左侧子树为所有比当且节点值小的子树，故在左侧递归搜索
  if (compare > 0) return searchBST(node.left, val)
  if (compare < 0) return searchBST(node.right, val)

  return node
}
