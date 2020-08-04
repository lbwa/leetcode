package medium199

import structures "leetcode-solutions/data-structures"

func rightSizeView(root *structures.BinaryTreeNode) []int {
	queue, answer := []*structures.BinaryTreeNode{root}, []int{}

	if root == nil {
		return answer
	}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 二叉树右视图，即仅保存每层节点的最后一个节点
			if i == levelSize-1 {
				answer = append(answer, current.Val.(int))
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return answer
}
