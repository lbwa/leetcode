import { maxDepth, TreeNode } from './index'

describe('N 叉树最大深度', () => {
  it('Should return 3', () => {
    const root = new TreeNode(1, [
      new TreeNode(3, [new TreeNode(5, []), new TreeNode(6, [])]),
      new TreeNode(2, []),
      new TreeNode(4, [])
    ])
    expect(maxDepth(root)).toBe(3)
  })
})
