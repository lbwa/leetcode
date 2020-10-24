import { BinaryTreeNode } from 'data-structures/binary-tree'
import { pathSum } from './solution'

describe('medium 113', () => {
  it('should return correct path', () => {
    expect(
      pathSum(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(
            4,
            new BinaryTreeNode(11, new BinaryTreeNode(7), new BinaryTreeNode(2))
          ),
          new BinaryTreeNode(
            8,
            new BinaryTreeNode(13),
            new BinaryTreeNode(4, new BinaryTreeNode(5), new BinaryTreeNode(1))
          )
        ),
        22
      )
    ).toEqual([
      [5, 4, 11, 2],
      [5, 8, 4, 5]
    ])
  })
})
