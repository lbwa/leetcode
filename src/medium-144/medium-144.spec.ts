import { preorderTraversal } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('144 二叉树的前序遍历', () => {
  it('Should return a pre-order list', () => {
    expect(
      preorderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 2, 3])
  })
})
