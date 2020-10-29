package medium129

import (
	structures "leetcode-solutions/data-structures"
	"strconv"
)

type node = structures.BinaryTreeNode

/*
	时间复杂度：O(N2)
	空间复杂度：O(N)
*/
func sumNumbers(root *node) int {
	paths := []string{}
	current := ""
	var dfs func(*node)
	dfs = func(root *node) {
		if root == nil {
			return
		}

		current += strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil {
			paths = append(paths, current)
		}
		dfs(root.Left)
		dfs(root.Right)
		current = current[:len(current)-1]
	}
	dfs(root)
	var ans int
	for _, path := range paths {
		if num, err := strconv.Atoi(path); err == nil {
			ans += num
		}
	}
	return ans
}
