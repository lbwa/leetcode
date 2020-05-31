import { maxDepth } from './index'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('104 二叉树的最大深度', () => {
  it('Should return max depth 3', () => {
    const node = new BinaryTreeNode(
      3,
      new BinaryTreeNode(9),
      new BinaryTreeNode(20, new BinaryTreeNode(15), new BinaryTreeNode(7))
    )
    expect(maxDepth(node)).toBe(3)
  })
})
