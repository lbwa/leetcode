import { LinkedListNode } from 'data-structures/linked-list'

export function reverseBetween(
  head: LinkedListNode | null,
  m: number, // 1 ≤ m ≤ n ≤ 链表长度
  n: number // 1 ≤ m ≤ n ≤ 链表长度
) {
  if (head === null) return head
  const sentinel = new LinkedListNode(Infinity, head)
  let prev: LinkedListNode = sentinel

  for (let i = 0; i < m - 1; i++) {
    prev = prev.next!
  }

  // 此时 prev.next 为需要被反转的首项，即第 m 项
  let current: LinkedListNode | null = prev.next

  // m = 2 时，第一次迭代如下
  // 12345 // next = current.next
  // 1245, next 3 -> 4 // current.next = next.next
  // 1245, next 3 -> 2 // next.next = prev.next
  // 13245 // prev.next = next
  for (let j = m; j < n; j++) {
    // 保存即将被反转的项
    const beReversed = current!.next!
    // 将当前待反转项分离出链表，此时该项仍保持对原 next 节点的链接
    current!.next = beReversed!.next // current!.next!.next
    // 反转待反转链接的 next 指向上一节点
    beReversed!.next = prev.next
    // 将反转后的项重新加入到链表中，完成一次反转
    prev.next = beReversed
  }

  return sentinel.next
}
