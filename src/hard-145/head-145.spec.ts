import { postorderTraversal } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('145 二叉树的后续遍历', () => {
  it('Should return a value list', () => {
    expect(
      postorderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3))
        )
      )
    ).toEqual([3, 2, 1])
  })
})
