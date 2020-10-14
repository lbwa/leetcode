import { BinaryTreeNode } from 'data-structures/binary-tree'
import { binaryTreePaths } from '.'

describe('easy 257 binary tree paths', () => {
  it('should pass', () => {
    expect(
      binaryTreePaths(
        new BinaryTreeNode(
          1,
          new BinaryTreeNode(2, null, new BinaryTreeNode(5)),
          new BinaryTreeNode(3)
        )
      )
    ).toEqual(['1->2->5', '1->3'])
  })
})
