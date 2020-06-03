import { BinaryTreeNode } from 'data-structures/binary-tree'

export function inorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer: number[] = []
  traversal(node, answer)
  return answer
}

function traversal(node: BinaryTreeNode<number> | null, answer: number[]) {
  if (node === null) return

  traversal(node.left, answer)
  answer.push(node.value)
  traversal(node.right, answer)
}
