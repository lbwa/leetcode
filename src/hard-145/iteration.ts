import { BinaryTreeNode } from 'data-structures/binary-tree'

export function postorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer: number[] = []
  const stack: (BinaryTreeNode<number> | null)[] = []
  if (node !== null) {
    stack.push(node)
  }

  while (stack.length > 0) {
    const top = stack.pop() as BinaryTreeNode<number> | null

    if (top !== null) {
      // 按照后序遍历的相反顺序压入栈中，使得在出栈时，各个节点是按照后序遍历的顺序出栈
      stack.push(top) // top 将遍历每个节点
      // 标识每个根节点的结尾，若当前 top 节点不再有子节点后，那么最后一个栈帧为 null
      stack.push(null)

      /**
       * 若存在至少一项左右子节点，那么将会继续进入下一轮迭代，直到当前节点不再具有
       * 子节点。此时，最后一个栈帧为 null，那么下一轮迭代将进入另一个 if 分支，并
       * 出栈栈顶形成后序遍历，并保存结果
       */
      if (top.right) {
        stack.push(top.right)
      }
      if (top.left) {
        stack.push(top.left)
      }
    } else {
      // 当栈顶帧为 null 时，表示已经深入到树的最底层，不再具有子节点，那么出栈栈顶
      // 形成中序遍历
      answer.push((stack.pop() as BinaryTreeNode<number>).value)
    }
  }

  return answer
}
