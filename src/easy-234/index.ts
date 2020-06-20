import { LinkedListNode } from 'data-structures/linked-list'

/**
 * 因为每个节点迭代一次，故时间复杂度为 O(N)，没有借助递归或额外数据结构，那么空间复杂
 * 度为 O(1)
 *
 * @keys
 * 1. 双指针找中点
 * 2. 反转中点（不含中点）至尾端的链表部分
 * 3. 迭代尾部链表部分，并与另外一边链表部分比较，得出是否为回文链表
 */
export function isPalindrome(head: LinkedListNode | null): boolean {
  if (head === null) return true

  let slow: LinkedListNode | null = head,
    fast: LinkedListNode | null = head

  /**
   * 结束迭代时，当 fast 为 null 时，表示当前链表项为偶数项，此时 slow 为链表的中项中
   * 的一项；当 fast.next 为 null 时，表示当前链表项为奇数项，此时 slow 为链表项
   */
  while (fast && fast.next) {
    slow = (slow as LinkedListNode).next
    fast = fast.next.next
  }

  /**
   * 若 fast 非 null，即 slow 为链表中项的唯一项，那么偏移一项，以便于后续反转从
   * slow 至链表结尾
   *
   * 为什么要偏移一项？因为在 “回文” 判断中，中项为唯一项时，不参与双指针判断，或与自身
   * 判断
   */
  if (fast !== null) {
    slow = (slow as LinkedListNode).next
  }

  let left: LinkedListNode | null = head,
    right = reverseLinkedList(slow)

  // 以反转的链表部分为基准
  while (right !== null) {
    if (right.val !== (left as LinkedListNode).val) return false
    left = (left as LinkedListNode).next
    right = right.next
  }

  /**
   * 上文解法会修改原链表，可于此在 return 结果前，反转之前修改的后半部分链表，以
   * 恢复链表顺序
   */

  return true
}

function reverseLinkedList(head: LinkedListNode | null) {
  let prev: LinkedListNode | null = null
  let current: LinkedListNode | null = head

  while (current) {
    const next = current.next
    current.next = prev
    prev = current
    current = next
  }

  return prev
}

/**
 * 二叉树存在一种特例就是单向链表，那么可借助二叉树的递归的后续遍历来实现回文链表的检测
 * 后续遍历结合双指针策略分别指向头部和尾部的节点
 *
 * 因为遍历了每个节点一次，那么时间复杂度为 O(N)，因为通过额外的递归栈空间来保存之前的
 * 节点，那么空间复杂度为 O(N)
 */
// export function isPalindrome(head: LinkedListNode | null) {
//   let left: LinkedListNode | null = head

//   function traversal(right: LinkedListNode | null): boolean {
//     if (right === null) return true

//     let answer = traversal(right.next)
//     answer =
//       answer && (left as LinkedListNode).val === (right as LinkedListNode).val
//     left = (left as LinkedListNode).next
//     return answer
//   }

//   return traversal(head)
// }
