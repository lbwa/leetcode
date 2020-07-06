import { hasPathSum } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('112 路径总和', () => {
  it('Bread-first traversal', () => {
    // [5,4,8,11,null,13,4,7,2,null,null,null,1]
    expect(
      hasPathSum(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(
            4,
            new BinaryTreeNode(
              11,
              new BinaryTreeNode(7, null, null),
              new BinaryTreeNode(2, null, null)
            ),
            null
          ),
          new BinaryTreeNode(
            8,
            new BinaryTreeNode(13, null, null),
            new BinaryTreeNode(4, null, new BinaryTreeNode(1, null, null))
          )
        ),
        22
      )
    ).toBeTruthy()
  })
})
