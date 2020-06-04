import { postorderTraversal as recursivePostOrderTraversal } from './recursion'
import { postorderTraversal as iterativePostOrderTraversal } from './iteration'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('145 二叉树的后续遍历', () => {
  it('Should return a post-order list with recursive traversal', () => {
    expect(
      recursivePostOrderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3))
        )
      )
    ).toEqual([3, 2, 1])
  })

  it('Should return a post-order list with iterative traversal', () => {
    expect(
      iterativePostOrderTraversal(
        new BinaryTreeNode(
          1,
          null,
          new BinaryTreeNode(2, new BinaryTreeNode(3))
        )
      )
    ).toEqual([3, 2, 1])
  })
})
