import { insertIntoBST } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('701 二叉搜索树中的插入操作', () => {
  it('Should return a node', () => {
    expect(
      insertIntoBST(
        new BinaryTreeNode(
          4,
          new BinaryTreeNode(
            2,
            new BinaryTreeNode(1, null, null),
            new BinaryTreeNode(3, null, null)
          ),
          new BinaryTreeNode(7, null, null)
        ),
        5
      )
    ).toEqual(
      new BinaryTreeNode(
        4,
        new BinaryTreeNode(
          2,
          new BinaryTreeNode(1, null, null),
          new BinaryTreeNode(3, null, null)
        ),
        new BinaryTreeNode(7, new BinaryTreeNode(5, null, null), null)
      )
    )
  })
})
