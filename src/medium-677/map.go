package medium677

// MapSum defines a kind of trie tree node
type MapSum struct {
	val  int
	next map[rune]*MapSum
}

// Constructor is used to initialize a MapSum instance
func Constructor() MapSum {
	return MapSum{next: make(map[rune]*MapSum)}
}

// Insert is used to insert key-value pair
func (root *MapSum) Insert(key string, val int) {
	for _, char := range key {
		if _, ok := root.next[char]; !ok {
			children := Constructor()
			root.next[char] = &children
		}
		root = root.next[char]
	}
	root.val = val
}

// Sum is used to calculate all prefix string sum
func (root *MapSum) Sum(prefix string) int {
	for _, char := range prefix {
		if _, ok := root.next[char]; !ok {
			return 0
		}
		root = root.next[char]
	}
	// 此时 node 为以 prefix 为前缀的 trie tree，那么求 node 的所有子节点的 val 之和（通过 dfs）
	return dfs(root)
}

func dfs(s *MapSum) (res int) {
	res += s.val
	for _, v := range s.next {
		res += dfs(v)
	}
	return
}
