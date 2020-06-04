import { BinaryTreeNode } from 'data-structures/binary-tree'

export function preorderTraversal(node: BinaryTreeNode<number>) {
  const answer: number[] = []
  traversal(node, answer)
  return answer
}

function traversal(node: BinaryTreeNode<number> | null, answer: number[]) {
  if (node === null) return
  answer.push(node.value)
  traversal(node.left, answer)
  traversal(node.right, answer)
}
