import { isValidBST } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('98 验证二叉搜索树', () => {
  it('Should return true', () => {
    expect(
      isValidBST(
        new BinaryTreeNode(2, new BinaryTreeNode(1), new BinaryTreeNode(3))
      )
    ).toBeTruthy()

    expect(isValidBST(null)).toBeTruthy()

    expect(isValidBST(new BinaryTreeNode(0))).toBeTruthy()
  })

  it('Should return false', () => {
    expect(
      isValidBST(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(1),
          new BinaryTreeNode(4, new BinaryTreeNode(3), new BinaryTreeNode(6))
        )
      )
    ).toBeFalsy()
  })
})
