import { deleteDuplicated } from '.'
import { LinkedListNode } from 'data-structures/linked-list'

describe('83 删除排序链表中的重复元素', () => {
  it('Should return a correct list', () => {
    expect(
      deleteDuplicated(
        new LinkedListNode(
          1,
          new LinkedListNode(1, new LinkedListNode(2, null))
        )
      )
    ).toEqual(new LinkedListNode(1, new LinkedListNode(2, null)))
  })
})
