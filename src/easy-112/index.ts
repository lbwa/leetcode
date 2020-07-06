import { BinaryTreeNode } from 'data-structures/binary-tree'

export function hasPathSum(
  root: BinaryTreeNode<number> | null,
  sum: number
): boolean {
  const nodeQueue: BinaryTreeNode<number>[] = []
  const valueQueue: number[] = []

  if (root === null) return false

  nodeQueue.push(root)
  valueQueue.push(root.value)

  while (nodeQueue.length) {
    const levelSize = nodeQueue.length

    for (let i = 0; i < levelSize; i++) {
      const node = nodeQueue.pop()!
      const current = valueQueue.pop()!

      if (node.left === null && node.right === null && current === sum)
        return true

      if (node.left) {
        nodeQueue.push(node.left)
        valueQueue.push(current + node.value)
      }
      if (node.right) {
        nodeQueue.push(node.right)
        valueQueue.push(current + node.value)
      }
    }
  }

  return false
}
