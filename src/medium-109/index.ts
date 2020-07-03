import { LinkedListNode } from 'data-structures/linked-list'
import { BinaryTreeNode } from 'data-structures/binary-tree'

export function sortedListToBST<V>(
  head: LinkedListNode<V> | null
): BinaryTreeNode<V> | null {
  if (head === null) return null

  // find mid node start
  let prev: LinkedListNode<V> | null = null,
    slow: LinkedListNode<V> = head,
    fast: LinkedListNode<V> | null = head
  while (fast && fast.next) {
    prev = slow
    slow = slow.next as LinkedListNode<V>
    fast = fast.next.next
  }
  // find mid node end

  const mid = slow
  const root = new BinaryTreeNode(mid.val)

  // 链表仅有唯一一项时
  if (head === mid) return root

  // 链表不止一项时
  if (prev !== null) {
    prev.next = null
  }

  // 将链表依据中心项，分为两部分，构建 **平衡** 的二叉搜索树
  root.left = sortedListToBST(head)
  root.right = sortedListToBST(mid.next)
  return root
}
