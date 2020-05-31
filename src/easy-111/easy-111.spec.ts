import { BinaryTreeNode } from 'data-structures/binary-tree'
import { minDepth } from '.'

describe('111 二叉树的最小深度', () => {
  it('Should return min depth 2', () => {
    const node = new BinaryTreeNode(
      3,
      new BinaryTreeNode(9),
      new BinaryTreeNode(20, new BinaryTreeNode(15), new BinaryTreeNode(7))
    )
    expect(minDepth(node)).toBe(2)
  })

  it('Should return min depth 2', () => {
    const node = new BinaryTreeNode(1, new BinaryTreeNode(2))
    expect(minDepth(node)).toBe(2)
  })
})
