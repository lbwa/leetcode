package medium138

type node struct {
	Val    int
	Next   *node
	Random *node
}

func copyRandomList(head *node) *node {
	if head == nil {
		return head
	}

	nodeMap := make(map[*node]*node)

	current := head
	for current != nil {
		nodeMap[current] = &node{Val: current.Val}
		current = current.Next
	}

	current = head
	for current != nil {
		node := nodeMap[current]
		node.Next = nodeMap[current.Next]
		node.Random = nodeMap[current.Random]
		current = current.Next
	}

	return nodeMap[head]
}
