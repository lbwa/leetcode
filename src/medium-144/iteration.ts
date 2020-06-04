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

    if (isDef(topFrame)) {
      // 按照树的前序遍历的反顺序将节点压入栈中，以备在出栈时，按照前序遍历的顺序处理
      // 各个遍历的各个节点
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

      // 占位标识符
      stack.push(null)
    } else {
      // 此 if 分支表示之前的压栈已经完成

      // 出栈栈顶，此时的出栈顺序即是树的前序遍历顺序
      answer.push(stack.pop()?.value)
    }
  }

  return answer
}
