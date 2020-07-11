package solution

import (
	structures "leetcode-solutions/data-structures"
)

// IterativeInOrderTraversal is medium 94 solution
func IterativeInOrderTraversal(root *structures.BinaryTreeNode) []int {
	stack := []*structures.BinaryTreeNode{}
	answer := []int{}

	if root == nil {
		return answer
	}

	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		latest := len(stack) - 1
		current = stack[latest]
		stack = stack[:latest]
		answer = append(answer, current.Val)
		current = current.Right
	}

	return answer
}

// func main() {
// 	treeA := structures.BinaryTreeNode{Val: 1, Right: &structures.BinaryTreeNode{Val: 2, Left: &structures.BinaryTreeNode{Val: 3}}}
// 	fmt.Println(IterativeInOrderTraversal(&treeA))
// }
