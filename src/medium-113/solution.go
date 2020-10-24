package medium113

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func pathSum(root *node, sum int) (answer [][]int) {
	path := []int{}
	var dfs func(*node, int)
	dfs = func(root *node, sum int) {
		if root == nil {
			return
		}
		sum -= root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && sum == 0 {
			answer = append(answer, append([]int(nil), path...))
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		path = path[:len(path)-1]
	}
	dfs(root, sum)
	return
}
