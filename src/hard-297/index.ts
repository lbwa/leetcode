import { BinaryTreeNode } from 'data-structures/binary-tree'

/**
 * @principles 基于前序遍历(DFS) 的二叉树序列化
 */
export function serialize<V>(root: BinaryTreeNode<V> | null) {
  const answer: (V | null)[] = []
  const stack: (BinaryTreeNode<V> | null)[] = []

  if (root === null) return JSON.stringify(answer)

  stack.push(root)

  while (stack.length) {
    const current = stack.pop()!
    answer.push(current ? current.value : null)

    // 当前节点不为 null 时，将其所有子节点加入栈中，以备后续子节点（含
    // null）的构建。
    if (current) {
      stack.push(current.right)
      stack.push(current.left)
    }
  }

  return JSON.stringify(answer)
}

export function deserialize<V>(chars: string) {
  const queue: (V | null)[] = JSON.parse(chars)
  return createTreeNode(queue)
}

function createTreeNode<V>(queue: (V | null)[]): BinaryTreeNode<V> | null {
  const val = queue.shift()!
  if (val === null || val === undefined) return null

  return new BinaryTreeNode(val, createTreeNode(queue), createTreeNode(queue))
}
