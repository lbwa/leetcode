import { NTreeNode } from 'data-structures/n-tree'

export function levelOrder<V>(root: NTreeNode<V> | null): V[][] {
  const queue: NTreeNode<V>[] = []
  const answer: V[][] = []

  if (root === null) return answer

  queue.push(root)

  while (queue.length) {
    answer.push([])
    const levelSize = queue.length

    for (let i = 0; i < levelSize; i++) {
      const current = queue.shift()!
      answer[answer.length - 1].push(current.value)
      current.children.forEach((child) => queue.push(child))
    }
  }

  return answer
}
