import { BinaryTreeNode } from 'data-structures/binary-tree'

export function isValidBST(node: BinaryTreeNode<number> | null) {
  const stack: (BinaryTreeNode<number> | null)[] = []
  let current: BinaryTreeNode<number> | null = node
  let prev: number = -Infinity

  while (current !== null || stack.length > 0) {
    // 依据中序遍历定义，左子树先压栈
    while (current !== null) {
      stack.push(current)
      current = current.left
    }
    current = stack.pop()!
    if (current.value <= prev) return false
    prev = current.value

    // 该行需要结合 line 8 的 while 条件来看，当 current 为最底层的最左节点时，那么
    // current.right 仍为 null, 但继续迭代，至 14 行，出栈得到父节点，即遍历父节点，
    // 并继续选取父节点的右节点，若存在，就继续遍历右子树
    current = current.right
  }

  return true
}
