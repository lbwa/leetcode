import { levelOrder } from '.'
import { NTreeNode } from 'data-structures/n-tree'

describe('429 N ary tree level order', () => {
  it('Should return a correct answer list', () => {
    expect(levelOrder(null)).toEqual([])
    expect(
      levelOrder(
        new NTreeNode(1, [
          new NTreeNode(3, [new NTreeNode(5), new NTreeNode(6)]),
          new NTreeNode(2),
          new NTreeNode(4)
        ])
      )
    ).toEqual([[1], [3, 2, 4], [5, 6]])
  })
})
