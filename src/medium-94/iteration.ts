import { BinaryTreeNode } from 'data-structures/binary-tree'
import { isDef } from 'src/shared'

export function inOrderTraversal(node: BinaryTreeNode<number> | null) {
  const stack: (BinaryTreeNode<number> | null)[] = []
  const answer: number[] = []
  let current: BinaryTreeNode<number> | null = node

  while (current !== null || stack.length > 0) {
    // 依据中序遍历定义，左子树先压栈
    while (current !== null) {
      stack.push(current)
      current = current.left
    }

    current = stack.pop()!
    answer.push(current.value)
    // 该行需要结合 line 11 的 while 条件来看，当 current 为最底层的最左节点时，那么
    // current.right 仍为 null, 但继续迭代，至 16 行，出栈得到父节点，即遍历父节点，
    // 并继续选取父节点的右节点，若存在，就继续遍历右子树
    current = current.right
  }

  return answer
}

export function inorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer: number[] = []
  const stack: (BinaryTreeNode<number> | null)[] = []

  if (node !== null) {
    stack.push(node)
  }

  while (stack.length > 0) {
    const topFrame = stack.pop() as BinaryTreeNode<number> | null

    /**
     * 按照中序遍历的相反顺序压入栈中，以备在出栈时，按照中序遍历的顺序处理节点
     * 中序遍历的原顺序为 左->中->右
     */
    if (isDef(topFrame)) {
      /**
       * 若存在至少一项左右子节点，那么将会继续进入下一轮迭代，直到当前节点不再具有
       * 子节点。此时，最后一个栈帧为 null，那么下一轮迭代将进入另一个 if 分支，并
       * 出栈栈顶形成中序遍历，并保存结果
       */
      if (topFrame.right) {
        stack.push(topFrame.right)
      }
      stack.push(topFrame)
      // 标识每个根节点的结尾，当当前节点不再具有子节点后，那么最后一个栈帧将为 null
      stack.push(null)
      /**
       * 若存在至少一项左右子节点，那么将会继续进入下一轮迭代，直到当前节点不再具有
       * 子节点。此时，最后一个栈帧为 null，那么下一轮迭代将进入另一个 if 分支，并
       * 出栈栈顶形成中序遍历，并保存结果
       */
      if (topFrame.left) {
        stack.push(topFrame.left)
      }
    } else {
      // 当栈顶帧为 null 时，表示已经深入到树的最底层，不再具有子节点，那么出栈栈顶
      // 形成中序遍历
      answer.push(stack.pop()?.value as number)
    }
  }

  return answer
}
