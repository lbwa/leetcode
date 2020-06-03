import { inorderTraversal } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('94 二叉树的中序遍历', () => {
  it('Should return a value list', () => {
    expect(
      inorderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 3, 2])
  })
})
