package medium889

import (
	structures "leetcode-solutions/data-structures"
)

type node = structures.BinaryTreeNode

func constructFromPrePost(pre, post []int) *node {
	if len(pre) < 1 || len(post) < 1 {
		return nil
	}
	val := pre[0]
	root := &node{Val: val}
	if len(pre) == 1 {
		return root
	}
	var i int
	for i = range post {
		if post[i] == pre[1] {
			i++
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:i+1], post[:i])
	root.Right = constructFromPrePost(pre[i+1:], post[i:len(post)-1])
	return root
}
