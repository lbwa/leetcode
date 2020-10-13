package easy257

import (
	structures "leetcode-solutions/data-structures"
	"strconv"
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
