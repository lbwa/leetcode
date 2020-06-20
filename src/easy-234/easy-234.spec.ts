import { isPalindrome } from '.'
import { LinkedListNode } from 'data-structures/linked-list'

describe('234 回文链表', () => {
  it('Should return a correct result', () => {
    expect(
      isPalindrome(
        new LinkedListNode(
          1,
          new LinkedListNode(
            2,
            new LinkedListNode(2, new LinkedListNode(1, null))
          )
        )
      )
    ).toBeTruthy()

    expect(
      isPalindrome(new LinkedListNode(1, new LinkedListNode(2, null)))
    ).toBeFalsy()
  })
})
