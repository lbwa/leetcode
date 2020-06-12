import { BinaryTreeNode } from 'data-structures/binary-tree'

export function levelOrder<V>(node: BinaryTreeNode<V> | null): V[][] {
  const queue: BinaryTreeNode<V>[] = []
  const answer: V[][] = []

  if (node === null) return answer

  queue.push(node)

  while (queue.length) {
    // 为当层遍历的结果预先分配存储空间
    answer.push([])

    // 每次取第 sn 层的 n 个节点进行迭代
    const levelSize = queue.length
    for (let i = 0; i < levelSize; i++) {
      const current = queue.shift()!
      answer[answer.length - 1].push(current.value)

      // 将第 sn 层节点的所有子节点按照遍历顺序逐一加入队列中，用于下一次 while 迭代
      if (current.left) queue.push(current.left)
      if (current.right) queue.push(current.right)
    }
    // 当次 while 迭代完成后，queue 中剩余的是 sn+1 层待遍历节点
  }

  return answer
}
