import { isValidBST as recursiveIsValidBST } from './recursion'
import { isValidBST as iterativeIsValidBST } from './iteration'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('98 验证二叉搜索树', () => {
  it('Should return true with recursion', () => {
    expect(
      recursiveIsValidBST(
        new BinaryTreeNode(2, new BinaryTreeNode(1), new BinaryTreeNode(3))
      )
    ).toBeTruthy()

    expect(recursiveIsValidBST(null)).toBeTruthy()

    expect(recursiveIsValidBST(new BinaryTreeNode(0))).toBeTruthy()
  })

  it('Should return true with iteration', () => {
    expect(
      iterativeIsValidBST(
        new BinaryTreeNode(2, new BinaryTreeNode(1), new BinaryTreeNode(3))
      )
    ).toBeTruthy()

    expect(iterativeIsValidBST(null)).toBeTruthy()

    expect(iterativeIsValidBST(new BinaryTreeNode(0))).toBeTruthy()
  })

  it('Should return false with recursion', () => {
    expect(
      recursiveIsValidBST(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(1),
          new BinaryTreeNode(4, new BinaryTreeNode(3), new BinaryTreeNode(6))
        )
      )
    ).toBeFalsy()

    expect(
      recursiveIsValidBST(new BinaryTreeNode(1, null, new BinaryTreeNode(1)))
    ).toBeFalsy()
  })

  it('Should return false with iteration', () => {
    expect(
      iterativeIsValidBST(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(1),
          new BinaryTreeNode(4, new BinaryTreeNode(3), new BinaryTreeNode(6))
        )
      )
    ).toBeFalsy()

    expect(
      iterativeIsValidBST(new BinaryTreeNode(1, null, new BinaryTreeNode(1)))
    ).toBeFalsy()
  })
})
