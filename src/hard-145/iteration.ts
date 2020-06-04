import { BinaryTreeNode } from 'data-structures/binary-tree'

export function postorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer: number[] = []
  const stack: (BinaryTreeNode<number> | null)[] = []
  if (node !== null) {
    stack.push(node)
  }

  while (stack.length > 0) {
    const top = stack.pop()

    if (top !== null && top !== undefined) {
      // 按照后序遍历的相反顺序压入栈中，使得在出栈时，各个节点是按照后序遍历的顺序出栈
      stack.push(top)
      stack.push(null)
      if (top.right) {
        stack.push(top.right)
      }
      if (top.left) {
        stack.push(top.left)
      }
    } else {
      answer.push((stack.pop() as BinaryTreeNode<number>).value)
    }
  }

  return answer
}
