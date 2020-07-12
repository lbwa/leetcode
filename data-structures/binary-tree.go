package structures

// BinaryTreeNode is used to create binary tree node
type BinaryTreeNode struct {
	/*
		An empty interface may hold values of any type.(Every type implements at
		least zero methods.)
		Empty interfaces are used by code that handles values of unknown type
		https://tour.golang.org/methods/14
		https://blog.golang.org/laws-of-reflection
		https://golang.org/ref/spec#Interface_types
	*/
	Val   interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}
