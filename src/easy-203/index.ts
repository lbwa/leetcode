import { LinkedListNode } from 'data-structures/linked-list'

export function removeElements(
  head: LinkedListNode | null,
  val: number
): LinkedListNode | null {
  if (head === null) return null

  // 哨兵节点，用于规避处理多种头节点删除的复杂情况
  const sentinel: LinkedListNode = new LinkedListNode(-1, head)
  let current: LinkedListNode | null = head
  let prev: LinkedListNode = sentinel

  while (current) {
    if (current.val === val) {
      // 发生匹配时，将在链表中通过以下语句删除当前节点
      prev.next = current.next
    } else {
      prev = current
    }

    current = current.next
  }

  return sentinel.next
}
