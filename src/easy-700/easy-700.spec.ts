import { searchBST } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('700 二叉搜索树中的搜索', () => {
  it('Should return null', () => {
    expect(
      searchBST(
        new BinaryTreeNode(
          2,
          new BinaryTreeNode(1, null, null),
          new BinaryTreeNode(3, null, null)
        ),
        5
      )
    ).toBeNull()
  })

  it('Should return a node', () => {
    expect(
      searchBST(
        new BinaryTreeNode(
          4,
          new BinaryTreeNode(
            2,
            new BinaryTreeNode(1, null, null),
            new BinaryTreeNode(3, null, null)
          ),
          new BinaryTreeNode(7, null, null)
        ),
        2
      )
    ).toEqual(
      new BinaryTreeNode(
        2,
        new BinaryTreeNode(1, null, null),
        new BinaryTreeNode(3, null, null)
      )
    )
  })
})
