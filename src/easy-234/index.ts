import { LinkedListNode } from 'data-structures/linked-list'

/**
 * 二叉树存在一种特例就是单向链表，那么可借助二叉树的递归的后续遍历来实现回文链表的检测
 * 后续遍历结合双指针策略分别指向头部和尾部的节点
 *
 * 因为遍历了每个节点一次，那么时间复杂度为 O(N)，因为通过额外的递归栈空间来保存之前的
 * 节点，那么空间复杂度为 O(N)
 */
export function isPalindrome(head: LinkedListNode | null) {
  let left: LinkedListNode | null = head

  function traversal(right: LinkedListNode | null): boolean {
    if (right === null) return true

    let answer = traversal(right.next)
    answer =
      answer && (left as LinkedListNode).val === (right as LinkedListNode).val
    left = (left as LinkedListNode).next
    return answer
  }

  return traversal(head)
}
