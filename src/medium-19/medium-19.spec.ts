import { removeNthFromEnd } from '.'
import { createLinkedList, LinkedListNode } from 'data-structures/linked-list'

describe('medium 19', () => {
  it('removeNthFromEnd([1], 1)', () => {
    expect(
      removeNthFromEnd(createLinkedList(1) as LinkedListNode, 1)
    ).toBeNull()
  })

  it('removeNthFromEnd([5, 4, 3, 2, 1], 2)', () => {
    expect(removeNthFromEnd(createLinkedList(5) as LinkedListNode, 2)).toEqual(
      new LinkedListNode(
        5,
        new LinkedListNode(
          4,
          new LinkedListNode(3, new LinkedListNode(1, null))
        )
      )
    )
  })

  it('removeNthFromEnd([5, 4, 3, 2, 1], 1)', () => {
    expect(removeNthFromEnd(createLinkedList(5) as LinkedListNode, 1)).toEqual(
      new LinkedListNode(
        5,
        new LinkedListNode(
          4,
          new LinkedListNode(3, new LinkedListNode(2, null))
        )
      )
    )
  })
})
