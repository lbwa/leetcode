import { deleteDuplicates } from '.'
import { LinkedListNode } from 'data-structures/linked-list'
import { serializeLinkedList } from 'src/shared'

describe('82 删除有序链表中的重复项 II', () => {
  it('Should return a correct list', () => {
    expect(
      serializeLinkedList(
        deleteDuplicates(
          new LinkedListNode(
            1,
            new LinkedListNode(
              2,
              new LinkedListNode(
                3,
                new LinkedListNode(
                  3,
                  new LinkedListNode(
                    4,
                    new LinkedListNode(4, new LinkedListNode(5, null))
                  )
                )
              )
            )
          )
        )
      )
    ).toEqual([1, 2, 5])
    expect(
      serializeLinkedList(
        deleteDuplicates(
          new LinkedListNode(
            1,
            new LinkedListNode(
              1,
              new LinkedListNode(
                1,
                new LinkedListNode(2, new LinkedListNode(3, null))
              )
            )
          )
        )
      )
    ).toEqual([2, 3])
  })
})
