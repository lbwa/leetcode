package medium117

// Node is a binary tree node structure
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	stack := []*Node{root}
	var prev *Node

	if root == nil {
		return root
	}

	for len(stack) > 0 {
		levelSize := len(stack)

		for i := 0; i < levelSize; i++ {
			current := stack[0]
			stack = stack[1:]

			if prev != nil {
				prev.Next = current
			}
			prev = current

			if current.Left != nil {
				stack = append(stack, current.Left)
			}
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
		}

		// 跨层不应链接
		prev = nil
	}

	return root
}
