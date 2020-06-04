import { preorderTraversal as recursivePreorderTraversal } from './recursion'
import { preorderTraversal as iterativePreorderTraversal } from './iteration'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('144 二叉树的前序遍历', () => {
  it('Should return a pre-order list with recursive pre-order traversal', () => {
    expect(
      recursivePreorderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 2, 3])
  })

  it('Should return a pre-order list with iterative pre-order traversal', () => {
    expect(
      iterativePreorderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3, null, null), null)
        )
      )
    ).toEqual([1, 2, 3])
  })
})
