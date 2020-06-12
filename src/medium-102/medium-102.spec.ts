import { levelOrder } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('102 二叉树的层次遍历（广度优先搜索 BFS）', () => {
  it('Should return a correct list', () => {
    expect(
      levelOrder(
        new BinaryTreeNode(
          3,
          new BinaryTreeNode(9),
          new BinaryTreeNode(20, new BinaryTreeNode(15), new BinaryTreeNode(7))
        )
      )
    ).toEqual([[3], [9, 20], [15, 7]])
  })
})
