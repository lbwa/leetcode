package medium129

import (
	structures "leetcode-solutions/data-structures"
	"strconv"
)

type node = structures.BinaryTreeNode

/*
	思路同 medium 113 路径总和 ii
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

/*
	时间复杂度：O(N)
	空间复杂度：O(N)
*/
func sumNumbers0(root *node) int {
	var dfs func(*node, int) int
	dfs = func(root *node, prev int) int {
		if root == nil {
			return 0
		}
		sum := prev*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}
