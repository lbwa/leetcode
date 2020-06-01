import { invertTree } from './index'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('226 翻转二叉树', () => {
  it('Should invert binary tree', () => {
    expect(invertTree(new BinaryTreeNode(0))).toEqual(new BinaryTreeNode(0))
    expect(
      invertTree(
        new BinaryTreeNode(
          0,
          new BinaryTreeNode(1, new BinaryTreeNode(3), new BinaryTreeNode(4)),
          new BinaryTreeNode(2, new BinaryTreeNode(5), new BinaryTreeNode(6))
        )
      )
    ).toEqual(
      new BinaryTreeNode(
        0,
        new BinaryTreeNode(2, new BinaryTreeNode(6), new BinaryTreeNode(5)),
        new BinaryTreeNode(1, new BinaryTreeNode(4), new BinaryTreeNode(3))
      )
    )
  })
})
