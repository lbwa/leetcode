package medium707

// ListNode defines a node from singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList defines a singly linked list with size
type LinkedList struct {
	size int
	head *ListNode
}

// CreateLinkedList is used to create a singly linked list
func CreateLinkedList() LinkedList {
	/*
		KEY: singly linked list should include a sentinel node for simplify head deletion
	*/
	return LinkedList{size: 0, head: &ListNode{}}
}

// Get is used to get a specific node
func (list *LinkedList) Get(index int) int {
	if index > (list.size - 1) {
		return -1
	}
	current := list.head.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

// AddAtHead is used to add a node from head
func (list *LinkedList) AddAtHead(val int) {
	current := &ListNode{Val: val, Next: list.head.Next}
	list.head.Next = current
	list.size++
}

// AddAtTail is used to add a node from tail
func (list *LinkedList) AddAtTail(val int) {
	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: val}
	list.size++
}

// AddAtIndex is used to add node based on index
func (list *LinkedList) AddAtIndex(index, val int) {
	if index == list.size {
		list.AddAtTail(val)
		return
	}
	// ignore oversize
	if index > list.size {
		return
	}
	if index < 0 {
		list.AddAtHead(val)
		return
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = &ListNode{Val: val, Next: current.Next}
	list.size++
}

// DeleteAtIndex is used to delete a node based on index
func (list *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || (index >= list.size) {
		return
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	list.size--
}
