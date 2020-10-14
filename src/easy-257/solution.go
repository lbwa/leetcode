package easy257

import (
	structures "leetcode-solutions/data-structures"
	"strconv"
	"strings"
)

type node = structures.BinaryTreeNode

func binaryTreePaths(root *node) (paths []string) {
	createPath(root, "", &paths)
	return
}

func createPath(root *node, path string, paths *[]string) {
	if root == nil {
		return
	}
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
	} else {
		path += "->"
		createPath(root.Left, path, paths)
		createPath(root.Right, path, paths)
	}
}

func binaryTreePaths0(root *node) (paths []string) {
	if root == nil {
		return
	}
	var dfs func(root *node, path []string)
	dfs = func(root *node, path []string) {
		if root == nil {
			return
		}
		path = append(path, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right == nil {
			paths = append(paths, strings.Join(path, "->"))
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, []string{})
	return
}
