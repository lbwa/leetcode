import { hasCycle } from '.'
import { LinkedListNode } from 'data-structures/linked-list'

describe('141 环形链表', () => {
  it('Should return false', () => {
    expect(
      hasCycle(new LinkedListNode(3, new LinkedListNode(2, null)))
    ).toBeFalsy()

    expect(hasCycle(new LinkedListNode(1, null))).toBeFalsy()
  })

  it('Should return true', () => {
    const a = new LinkedListNode(0, null)
    const b = new LinkedListNode(11, null)
    const c = new LinkedListNode(22, null)
    const d = new LinkedListNode(33, null)
    a.next = b
    b.next = c
    c.next = d
    d.next = b
    expect(hasCycle(a)).toBeTruthy()

    const e = new LinkedListNode(1, null)
    const f = new LinkedListNode(2, null)
    e.next = f
    f.next = e
    expect(hasCycle(e)).toBeTruthy()
  })
})
