import { BinaryTreeNode } from 'data-structures/binary-tree'
import { isDef } from 'src/shared'

export function preorderTraversal(node: BinaryTreeNode<number> | null) {
  const answer = []
  const stack: (BinaryTreeNode<number> | null)[] = []
  if (node !== null) {
    stack.push(node)
  }
  while (stack.length > 0) {
    const topFrame = stack.pop()

    /**
     * 按照树的前序遍历的反顺序将节点压入栈中，以备在出栈时，按照前序遍历的顺序处理
     * 各个遍历的各个节点
     */
    if (isDef(topFrame)) {
      /**
       * 若存在至少一项左右子节点，那么将会继续进入下一轮迭代，直到当前节点不再具有
       * 子节点。此时，最后一个栈帧为 null，那么下一轮迭代将进入另一个 if 分支，并
       * 出栈栈顶形成中序遍历，并保存结果
       */
      if (topFrame.right) {
        // 先压右子树至栈中
        stack.push(topFrame.right)
      }
      if (topFrame.left) {
        // 再压左子树至栈中
        stack.push(topFrame.left)
      }
      // 最后压当前根节点至栈中，使得出栈时，第一个处理的节点是根节点
      stack.push(topFrame)

      // 标识数中每个根节点的结尾，当当前节点不再具有子节点后，那么最后一个栈中为 null
      // 进而在下一轮遍历中，进入下一个 if 分支，进而触发出栈栈顶，进而形成前序遍历
      stack.push(null)
    } else {
      // 当栈顶帧为 null 时，表示已经深入到树的最底层，不再具有子节点，那么此时出栈栈
      // 顶，此时的出栈顺序即是树的前序遍历顺序
      answer.push(stack.pop()?.value)
    }
  }

  return answer
}
