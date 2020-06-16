import { serialize, deserialize } from '.'
import { BinaryTreeNode } from 'data-structures/binary-tree'

describe('hard 297 二叉树的序列化和反序列化', () => {
  it('Should serialize a binary tree', () => {
    expect(serialize(null)).toBe('[]')
    expect(
      serialize(
        new BinaryTreeNode(
          1,
          new BinaryTreeNode(2, new BinaryTreeNode(3), new BinaryTreeNode(4)),
          new BinaryTreeNode(5)
        )
      )
    ).toBe(`[1,2,3,null,null,4,null,null,5,null,null]`)
  })

  it('Should deserialize a binary tree', () => {
    expect(deserialize(serialize(null))).toBeNull()
    expect(
      deserialize(
        serialize(
          new BinaryTreeNode(
            1,
            new BinaryTreeNode(2, new BinaryTreeNode(3), new BinaryTreeNode(4)),
            new BinaryTreeNode(5)
          )
        )
      )
    ).toEqual(
      new BinaryTreeNode(
        1,
        new BinaryTreeNode(2, new BinaryTreeNode(3), new BinaryTreeNode(4)),
        new BinaryTreeNode(5)
      )
    )
  })
})
