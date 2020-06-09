import { BinaryTreeNode } from 'data-structures/binary-tree'

export function isValidBST(node: BinaryTreeNode<number> | null) {
  const stack: (BinaryTreeNode<number> | null)[] = []
  let current: BinaryTreeNode<number> | null = node
  let prev: number = -Infinity

  while (current !== null || stack.length > 0) {
    while (current !== null) {
      stack.push(current)
      current = current.left
    }
    current = stack.pop()!
    if (current.value <= prev) return false
    prev = current.value
    current = current.right
  }

  return true
}
