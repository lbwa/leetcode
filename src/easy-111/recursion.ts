import { BinaryTreeNode } from 'data-structures/binary-tree'
import { isDef } from 'src/shared'

export function minDepth<V>(node: BinaryTreeNode<V> | null): number {
  if (!isDef(node)) return 0
  if (!isDef(node.left) && !isDef(node.right)) return 1

  let answer = Number.MAX_VALUE // or Number.MAX_SAFE_INTEGER
  if (isDef(node.left)) {
    answer = Math.min(minDepth(node.left), answer)
  }
  if (isDef(node.right)) {
    answer = Math.min(minDepth(node.right), answer)
  }
  return answer + 1
}
