import { isSymmetric } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('101 对称二叉树', () => {
  it('Should return true', () => {
    expect(
      isSymmetric(
        new BinaryTreeNode(
          1,
          new BinaryTreeNode(2, new BinaryTreeNode(3), new BinaryTreeNode(4)),
          new BinaryTreeNode(2, new BinaryTreeNode(4), new BinaryTreeNode(3))
        )
      )
    ).toBeTruthy()
  })

  it('Should return false', () => {
    expect(
      isSymmetric(
        new BinaryTreeNode(
          1,
          new BinaryTreeNode(2, null, new BinaryTreeNode(3)),
          new BinaryTreeNode(2, null, new BinaryTreeNode(3))
        )
      )
    ).toBeFalsy()
  })
})
