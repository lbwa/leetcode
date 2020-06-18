import { LinkedListNode } from 'data-structures/linked-list'

/**
 * 原理同 26 删除有序数组中的重复项，始终将待删除（重复）节点向后靠，并在指针达到尾端后
 * 删除尾缀部分
 * 目标是为了维护 [0, ...slow] 为结果部分
 */
export function deleteDuplicated(head: LinkedListNode | null) {
  if (head === null) return head

  let slow = head,
    fast = slow.next
  while (fast) {
    if (slow.val !== fast.val) {
      slow.next = fast
      slow = slow.next
    }
    fast = fast.next
  }

  slow.next = null

  return head
}
