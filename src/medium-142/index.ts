import { LinkedListNode } from 'data-structures/linked-list'

export function detectCycle(head: LinkedListNode) {
  let slow: LinkedListNode | null = head
  let fast: LinkedListNode | null = head

  while (fast && fast.next !== null) {
    fast = fast.next.next
    slow = (slow as LinkedListNode).next
    if (fast === slow) break
  }

  // 无环
  if (fast !== slow || fast === null || fast.next === null) return null

  // 有环，通过指针碰撞找到环开始点
  slow = head
  while (slow !== fast) {
    slow = (slow as LinkedListNode).next
    fast = (fast as LinkedListNode).next
  }
  return slow
}
