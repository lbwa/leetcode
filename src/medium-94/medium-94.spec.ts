import { inorderTraversal as recursiveInOrderTraversal } from './recursion'
import {
  inorderTraversal as iterativeInOrderTraversal,
  inOrderTraversal
} from './iteration'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('94 二叉树的中序遍历', () => {
  it('Should return a in-order list with recursive traversal', () => {
    expect(
      recursiveInOrderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 3, 2])

    expect(
      inOrderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 3, 2])
  })

  it('Should return a in-order list with iterative traversal', () => {
    expect(
      iterativeInOrderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 3, 2])
  })
})
