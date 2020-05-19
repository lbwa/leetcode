export class LinkedListNode<E> {
  constructor(public val: E, public next: LinkedListNode<E> | null) {}
}

/**
 * 反转单向链表
 * 1. 迭代每一个链表节点
 * 2. 在迭代过程中始终首先保存当前节点的下一个节点，以用于后续迭代
 * 3. 将当前节点的 next 指向 prev 节点
 * 4. 将 2 保存的节点赋值为 current 节点，返回 1 步骤继续迭代链表
 */
export function reverseLinkedList<E>(head: LinkedListNode<E>) {
  let prev = null
  let current: LinkedListNode<E> | null = head
  while (current) {
    // 临时保存 next 节点，以用于在完成交换后，恢复迭代
    const temp: LinkedListNode<E> | null = current.next
    current.next = prev
    prev = current
    current = temp
  }
  return prev
}
