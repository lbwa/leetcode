import { removeElements } from '.'
import { LinkedListNode } from 'data-structures/linked-list'

describe('203 移除链表元素', () => {
  it('Should remove head node', () => {
    const linkedList = new LinkedListNode(1, null)
    expect(removeElements(linkedList, 1)).toBeNull()
  })

  it('Should remove tail node', () => {
    const linkedList = new LinkedListNode(
      1,
      new LinkedListNode(2, new LinkedListNode(3, null))
    )
    expect(removeElements(linkedList, 3)).toEqual(
      new LinkedListNode(1, new LinkedListNode(2, null))
    )
  })

  it('Should remove multiple head part node', () => {
    const linkedList = new LinkedListNode(
      2,
      new LinkedListNode(2, new LinkedListNode(3, null))
    )
    expect(removeElements(linkedList, 2)).toEqual(new LinkedListNode(3, null))
  })
})
