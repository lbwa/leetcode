import { kthSmallest } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('230 二叉搜索树中第 K 小的元素', () => {
  it("Should return a node's value", () => {
    // [5,3,6,2,4,null,null,1]
    expect(
      kthSmallest(
        new BinaryTreeNode(
          5,
          new BinaryTreeNode(
            3,
            new BinaryTreeNode(2, new BinaryTreeNode(1, null, null), null),
            new BinaryTreeNode(4, null, null)
          ),
          new BinaryTreeNode(6, null, null)
        ),
        3
      )
    ).toBe(3)

    // [3,1,4,null,2]
    expect(
      kthSmallest(
        new BinaryTreeNode(
          3,
          new BinaryTreeNode(1, null, new BinaryTreeNode(2, null, null)),
          new BinaryTreeNode(4, null, null)
        ),
        1
      )
    ).toBe(1)
  })
})
