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
		// 深入 root 的左子树，进而通过下一层检查来检查是否存在符合要求的路径
		dfs(root.Left, sum)
		// 深入 root 的右子树，进而通过下一层检查来检查是否存在符合要求的路径
		dfs(root.Right, sum)
		/*
			移除 root 本身，因为包含 root 的路径已经在上述代码中遍历完成，故 **针对递归栈中** 之后的遍历得到路径中无 root 节点
		*/
		path = path[:len(path)-1]
	}
	// 前序遍历
	dfs(root, sum)
	return
}
