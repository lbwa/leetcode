import { bstFromPreOrder } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('1008 先序遍历构造二叉搜索树', () => {
  it('Should return a binary search tree', () => {
    expect(bstFromPreOrder([8, 5, 1, 7, 10, 12])).toEqual(
      new BinaryTreeNode(
        8,
        new BinaryTreeNode(5, new BinaryTreeNode(1), new BinaryTreeNode(7)),
        new BinaryTreeNode(10, null, new BinaryTreeNode(12))
      )
    )
  })
})
