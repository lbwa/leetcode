package medium208

// TrieViaArray defines a trie tree node based on array/slice
type TrieViaArray struct {
	isWord bool
	next   [26]*TrieViaArray
}

// Create is used to create a trie tree node
func Create() TrieViaArray {
	return TrieViaArray{}
}

// Insert is used to insert node if the path does not exist
func (root *TrieViaArray) Insert(word string) {
	for _, char := range word {
		index := char - 'a'
		if ptr := root.next[index]; ptr == nil {
			root.next[index] = new(TrieViaArray)
		}
		root = root.next[index]
	}
	root.isWord = true
}

// Search is used to search node from the current tree
func (root *TrieViaArray) Search(word string) bool {
	for _, char := range word {
		index := char - 'a'
		if ptr := root.next[index]; ptr == nil {
			return false
		}
		root = root.next[index]
	}
	return root.isWord
}

// StartsWith is used to judge prefix string
func (root *TrieViaArray) StartsWith(prefix string) bool {
	for _, char := range prefix {
		index := char - 'a'
		if ptr := root.next[index]; ptr == nil {
			return false
		}
		root = root.next[index]
	}
	return true
}
