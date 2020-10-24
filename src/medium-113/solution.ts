import { BinaryTreeNode } from 'data-structures/binary-tree'

export function pathSum(
  root: BinaryTreeNode<number> | null,
  sum: number
): number[][] {
  const path: number[] = []
  const answer: number[][] = []

  function dfs(root: BinaryTreeNode<number> | null, sum: number) {
    if (!root) {
      return
    }
    path.push(root.value)
    sum -= root.value
    if (!root.left && !root.right && sum === 0) {
      answer.push([...path])
    }
    dfs(root.left, sum)
    dfs(root.right, sum)
    path.pop()
  }

  dfs(root, sum)
  return answer
}
