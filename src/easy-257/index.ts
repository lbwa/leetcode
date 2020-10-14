import { BinaryTreeNode } from 'data-structures/binary-tree'

export function binaryTreePaths(root: BinaryTreeNode<number> | null): string[] {
  const ans: string[] = []
  function dfs(root: BinaryTreeNode<number> | null, paths: string[]) {
    if (!root) {
      return
    }
    paths.push(`${root.value}`)
    if (!root.left && !root.right) {
      ans.push(paths.join('->'))
    }
    dfs(root.left, [...paths])
    dfs(root.right, [...paths])
  }
  dfs(root, [])
  return ans
}
