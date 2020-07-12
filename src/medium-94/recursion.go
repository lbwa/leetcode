package solution

import structures "leetcode-solutions/data-structures"

// RecursiveInOrderTraversal is medium 94 solution
func RecursiveInOrderTraversal(root *structures.BinaryTreeNode) []int {
	answer := []int{}
	traversal(root, &answer)
	return answer
}

func traversal(root *structures.BinaryTreeNode, answer *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, answer)
	// type assertion
	// https://tour.golang.org/methods/15
	*answer = append(*answer, root.Val.(int))
	traversal(root.Right, answer)
}
