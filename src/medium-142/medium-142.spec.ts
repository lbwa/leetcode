import { detectCycle } from '.'
import { LinkedListNode } from 'data-structures/linked-list'

describe('142 环形链表 II', () => {
  it('Should return false', () => {
    expect(
      detectCycle(new LinkedListNode(3, new LinkedListNode(2, null)))
    ).toBeNull()
    expect(detectCycle(new LinkedListNode(1, null))).toBeNull()
  })

  it('Should return cycle position', () => {
    const a = new LinkedListNode(0, null)
    const b = new LinkedListNode(11, null)
    const c = new LinkedListNode(22, null)
    const d = new LinkedListNode(33, null)
    a.next = b
    b.next = c
    c.next = d
    d.next = b
    expect(detectCycle(a)).toBe(b)

    const e = new LinkedListNode(1, null)
    const f = new LinkedListNode(2, null)
    e.next = f
    f.next = e
    expect(detectCycle(e)).toBe(e)
  })
})
