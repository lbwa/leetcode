import { LinkedListNode } from 'data-structures/linked-list'

export function hasCycle(head: LinkedListNode) {
  let slow: LinkedListNode | null = head
  let fast: LinkedListNode | null = head

  while (fast && fast.next !== null) {
    fast = fast.next.next
    slow = (slow as LinkedListNode).next
    if (fast === slow) return true
  }

  return false
}
