package medium1609

import (
	structures "leetcode-solutions/data-structures"
	"math"
)

type node = structures.BinaryTreeNode

func isEvenOddTree(root *node) bool {
	if root == nil {
		return false
	}
	queue := []*node{root}
	levelIndex := -1
	for len(queue) > 0 {
		levelIndex++
		levelSize := len(queue)
		var prev int
		if isOdd(levelIndex) {
			prev = math.MaxInt64
		} else {
			prev = math.MinInt64
		}
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if isOdd(levelIndex) {
				if isOdd(current.Val) || (current.Val >= prev) {
					return false
				}
			} else {
				if !isOdd(current.Val) || (current.Val <= prev) {
					return false
				}
			}

			prev = current.Val

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return true
}

func isOdd(val int) bool {
	return val%2 > 0
}
