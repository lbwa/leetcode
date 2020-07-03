import { sortedListToBST } from '.'
import { LinkedListNode } from 'data-structures/linked-list'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('109 将有序链表转换为二叉搜索树', () => {
  it('Should return correct result', () => {
    expect(
      sortedListToBST(
        new LinkedListNode(
          -10,
          new LinkedListNode(
            -3,
            new LinkedListNode(
              0,
              new LinkedListNode(5, new LinkedListNode(9, null))
            )
          )
        )
      )
    ).toEqual(
      new BinaryTreeNode(
        0,
        new BinaryTreeNode(-3, new BinaryTreeNode(-10)),
        new BinaryTreeNode(9, new BinaryTreeNode(5))
      )
    )
  })
})
