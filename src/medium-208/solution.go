package medium208

// Trie defines a trie tree node
type Trie struct {
	isWord bool
	child  map[rune]*Trie
}

// CreateTrie is used to create a trie tree node
func CreateTrie() Trie {
	return Trie{isWord: false, child: make(map[rune]*Trie)}
}

// Insert is used to insert node if the path does not exist
func (root *Trie) Insert(word string) {
	node := root
	for _, char := range word {
		if _, ok := node.child[char]; !ok {
			child := CreateTrie()
			node.child[char] = &child
		}
		node = node.child[char]
	}
	// leaf node is word type node
	node.isWord = true
}

// Search is used to search node from the current tree
func (root *Trie) Search(word string) bool {
	node := root
	for _, char := range word {
		if _, ok := node.child[char]; !ok {
			return false
		}
		node = node.child[char]
	}
	return node.isWord
}

// StartsWith is used to judge prefix string
func (root *Trie) StartsWith(prefix string) bool {
	node := root
	for _, char := range prefix {
		if _, ok := node.child[char]; !ok {
			return false
		}
		node = node.child[char]
	}
	return true
}
