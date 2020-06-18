export class LinkedListNode<T = unknown> {
  constructor(public val: T, public next: LinkedListNode<T> | null) {}
}

export function createLinkedList(max: number): LinkedListNode<number> | null {
  if (max === 0) return null
  return new LinkedListNode<number>(max, createLinkedList(--max))
}
