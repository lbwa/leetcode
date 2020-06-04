import { BinaryTreeNode } from 'data-structures/binary-tree'
import { isDef } from 'src/shared'

export function inorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer: number[] = []
  const stack: (BinaryTreeNode<number> | null)[] = []

  if (node !== null) {
    stack.push(node)
  }

  while (stack.length > 0) {
    const topFrame = stack.pop()

    if (isDef(topFrame)) {
      /**
       * 按照中序遍历的相反顺序压入栈中，以备在出栈时，按照中序遍历的顺序处理节点
       * 中序遍历的原顺序为 左->中->右
       */
      if (topFrame.right) {
        stack.push(topFrame.right)
      }
      stack.push(topFrame)
      stack.push(null)
      if (topFrame.left) {
        stack.push(topFrame.left)
      }
    } else {
      // 当栈顶帧为 null 时，表示之前压栈已经完成，那么出栈栈顶，并保存该值
      answer.push(stack.pop()?.value as number)
    }
  }

  return answer
}
