import { sortedArrayToBST } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('108 将有序数组转换为二叉搜索树', () => {
  it('should return correct results', () => {
    expect(sortedArrayToBST([])).toBeNull()
    expect(sortedArrayToBST([-10, -3, 0, 5, 9])).toEqual(
      new BinaryTreeNode(
        0,
        new BinaryTreeNode(-10, null, new BinaryTreeNode(-3)),
        new BinaryTreeNode(5, null, new BinaryTreeNode(9))
      )
    )
    expect(sortedArrayToBST([0, 1, 2, 3, 4, 5])).toEqual(
      new BinaryTreeNode(
        2,
        new BinaryTreeNode(0, null, new BinaryTreeNode(1)),
        new BinaryTreeNode(4, new BinaryTreeNode(3), new BinaryTreeNode(5))
      )
    )
  })
})
