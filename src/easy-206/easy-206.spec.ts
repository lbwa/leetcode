import { reverseLinkedList as iterativeReverseLinkedList } from './iteration'
import { reverseLinkedList as recursiveReverseLinkedList } from './recursion'
import { LinkedListNode } from 'data-structures/linked-list'

function createLinkedList(max: number): LinkedListNode<number> {
  if (max === 0) return new LinkedListNode(0, null)
  return new LinkedListNode(max, createLinkedList(--max))
}

describe('Easy 206 reverse linked-list', () => {
  it('Should be reversed', () => {
    const head = createLinkedList(5)
    const resultHead = iterativeReverseLinkedList(head)
    let current = resultHead
    while (current) {
      if (current.next) {
        expect(current.val).toBeLessThan(current.next.val)
      }
      current = current.next
    }
  })

  it('Should be reversed', () => {
    const head = createLinkedList(5)
    const resultHead = recursiveReverseLinkedList(head)
    let current = resultHead
    while (current) {
      if (current.next) {
        expect(current.val).toBeLessThan(current.next.val)
      }
      current = current.next
    }
  })
})
