package easy108

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func sortedArrayToBST(nums []int) *node {
	var createTree func(nums []int, low int, high int) *node
	createTree = func(nums []int, low int, high int) *node {
		if low > high {
			return nil
		}
		index := low + ((high - low) / 2)
		return &node{
			Val:   nums[index],
			Left:  createTree(nums, low, index-1),
			Right: createTree(nums, index+1, high),
		}
	}
	return createTree(nums, 0, len(nums)-1)
}
