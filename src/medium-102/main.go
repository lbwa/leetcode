package main

import "fmt"

// TreeNode is binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeA := TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	fmt.Println(levelOrder(&treeA))
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	answer := [][]int{}

	if root == nil {
		return answer
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		answer = append(answer, []int{})

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			levelIndex := len(answer) - 1
			answer[levelIndex] = append(answer[levelIndex], current.Val)

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
