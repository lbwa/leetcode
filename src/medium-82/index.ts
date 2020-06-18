import { LinkedListNode } from 'data-structures/linked-list'

export function deleteDuplicates(head: LinkedListNode | null) {
  const sentinel = new LinkedListNode(NaN, head)
  let current = sentinel
  while (current.next && current.next.next) {
    if (current.next.val === current.next.next.val) {
      let temp = current.next
      while (temp.next && temp.val === temp.next.val) {
        temp = temp.next
      }
      current.next = temp.next
    } else {
      current = current.next
    }
  }
  return sentinel.next
}
