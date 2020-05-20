import { LinkedListNode } from '../../data-structures/linked-list'

/**
 * 删除给定链表的倒数第 k 个节点
 */
export function removeNthFromEnd(head: LinkedListNode, number: number) {
  let low: LinkedListNode | null = head
  let fast: LinkedListNode | null = head
  let prev: LinkedListNode | null = null
  let step = number
  // 快指针先走 number 步
  while (step-- && fast) {
    fast = fast.next
  }
  // 快慢指针同时并进，直到快指针达到链表的 null 节点，此时 low 节点即为待删除节点
  while (fast) {
    prev = low
    fast = fast.next
    low = (low as LinkedListNode).next
  }

  // 此时 prev 为 null 则表示待删除节点的上一个节点为 null，那么即为 head 节点待删除
  if (prev === null) {
    return head.next
  }

  // low 慢指针对应了待删除节点，删除之
  prev.next = (low as LinkedListNode).next
  return head
}
