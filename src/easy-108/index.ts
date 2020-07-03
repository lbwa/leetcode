import { BinaryTreeNode } from 'data-structures/binary-tree'

export function sortedArrayToBST<V>(nums: V[]): BinaryTreeNode<V> | null {
  return createTreeNode(nums, 0, nums.length - 1)
}

function createTreeNode<V>(nums: V[], low: number, high: number) {
  if (low > high) return null

  const pivotIndex = low + ((high - low) >> 1)
  const rootNode = new BinaryTreeNode(nums[pivotIndex])
  rootNode.left = createTreeNode(nums, low, pivotIndex - 1)
  rootNode.right = createTreeNode(nums, pivotIndex + 1, high)
  return rootNode
}
