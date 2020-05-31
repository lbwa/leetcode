import { isDef } from 'src/shared'

export class TreeNode<V> {
  constructor(public value: V, public children: TreeNode<V>[]) {}
}

export function maxDepth<V>(node: TreeNode<V>): number {
  if (!isDef(node)) return 0
  return (
    node.children.reduce(
      (depth, child) => Math.max(maxDepth(child), depth),
      0
    ) + 1
  )
}
