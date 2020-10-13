package easy783

import (
	structures "leetcode-solutions/data-structures"
	"math"
)

type node = structures.BinaryTreeNode

func minDiffInBST(root *node) int {
	ans, prev := math.MaxInt64, -1
	var dfs func(root *node)
	dfs = func(root *node) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != -1 && (root.Val.(int)-prev < ans) {
			ans = root.Val.(int) - prev
		}
		prev = root.Val.(int)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
