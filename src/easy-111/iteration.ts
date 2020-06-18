import { BinaryTreeNode } from 'data-structures/binary-tree'

/**
 * @key 借助 二叉树层序遍历 的框架实现最小深度探测
 */
export function minDepth<V>(root: BinaryTreeNode<V> | null) {
  const queue: BinaryTreeNode<V>[] = []
  let depth = 0

  if (root === null) return depth

  queue.push(root)

  while (queue.length) {
    let levelSize = queue.length

    depth++

    // 开始遍历第 depth(始于 1) 层节点，根节点属于第 1 层
    for (let i = 0; i < levelSize; i++) {
      const current = queue.shift()!

      // 当前 current 为叶子节点（即所有子节点为 null）时，返回当前深度
      if (current.left === null && current.right === null) return depth

      if (current.left) queue.push(current.left)
      if (current.right) queue.push(current.right)
    }
  }

  return depth
}
