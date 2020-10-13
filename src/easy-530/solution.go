package easy530

import (
	structures "leetcode-solutions/data-structures"
	"math"
)

type node = structures.BinaryTreeNode

func getMinimumDifference(root *node) int {
	answer, prev := math.MaxInt64, -1
	var dfs func(root *node)
	dfs = func(root *node) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != -1 && (root.Val-prev < answer) {
			answer = root.Val - prev
		}
		prev = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return answer
}
