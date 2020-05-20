export class LinkedListNode {
  constructor(public val: unknown, public next: LinkedListNode | null) {}
}

export function createLinkedList(max: number): LinkedListNode | null {
  if (max === 0) return null
  return new LinkedListNode(max, createLinkedList(--max))
}
