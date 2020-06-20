import { LinkedListNode } from 'data-structures/linked-list'

// 反转链表的 [m, n] 项节点
export function reverseBetween(
  head: LinkedListNode | null,
  m: number,
  n: number
): LinkedListNode | null {
  if (head === null) return head

  if (m === 1) {
    // 相当于反转前 n 个项
    return reverseUntil(head, n)
  }

  // 在原始链表中，头部为索引 1，需要反转的头部项为 m，通过递归规律，那么 head.next
  // 的需要反转的头部项为 m - 1，那么 head.next.next 的需要反转的头部项为 m - 1 - 1
  head.next = reverseBetween(head.next, m - 1, n - 1)

  return head
}

/**
 * 反转链表的前 n 个节点
 * @key 本质上是基于 easy 206 中递归反转链表，不同的是，记住了最后一次反转时的下一节点
 * 即未反转部分的头节点，以用于反转部分的尾节点的 `next` 链接
 * @see https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/di-gui-fan-zhuan-lian-biao-de-yi-bu-fen#er-fan-zhuan-lian-biao-qian-n-ge-jie-dian
 */
function reverseUntil(head: LinkedListNode, n: number) {
  // 该变量保持对反转后的链表尾项应该链接的项，即未发生反转部分的首项
  let successor: LinkedListNode | null = null

  function reverse(head: LinkedListNode, n: number): LinkedListNode {
    if (n === 1) {
      successor = head.next
      // 反转后 head 为反转部分的末项
      return head
    }

    const last = reverse(head.next as LinkedListNode, n - 1)

    ;(head.next as LinkedListNode).next = head
    head.next = successor

    return last
  }

  return reverse(head, n)
}
