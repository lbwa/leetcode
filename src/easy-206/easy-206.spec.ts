import { LinkedListNode, reverseLinkedList } from './index'

function createLinkedList(max: number): LinkedListNode<number> {
  if (max === 0) return new LinkedListNode(0, null)
  return new LinkedListNode(max, createLinkedList(--max))
}

describe('Easy 206 reverse linked-list', () => {
  it('Should be reversed', () => {
    const head = createLinkedList(5)
    const resultHead = reverseLinkedList(head)
    let current = resultHead
    while (current) {
      if (current.next) {
        expect(current.val).toBeLessThan(current.next.val)
      }
      current = current.next
    }
  })
})
