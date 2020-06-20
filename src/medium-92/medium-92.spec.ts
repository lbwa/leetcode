import { reverseBetween as recursiveReverseBetween } from './recursion'
import { LinkedListNode } from 'data-structures/linked-list'

describe('92 反转链表 II', () => {
  it('Should reverse items between m and n', () => {
    expect(
      recursiveReverseBetween(
        new LinkedListNode(
          1,
          new LinkedListNode(
            2,
            new LinkedListNode(
              3,
              new LinkedListNode(4, new LinkedListNode(5, null))
            )
          )
        ),
        2,
        4
      )
    ).toEqual(
      new LinkedListNode(
        1,
        new LinkedListNode(
          4,
          new LinkedListNode(
            3,
            new LinkedListNode(2, new LinkedListNode(5, null))
          )
        )
      )
    )
  })
})
