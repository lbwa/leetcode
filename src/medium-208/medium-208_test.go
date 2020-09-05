package medium208

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := CreateTrie()
	tree.Insert("apple")
	shared.Expect(t, tree.Search("apple"), true)
	shared.Expect(t, tree.Search("app"), false)
	shared.Expect(t, tree.StartsWith("app"), true)
	tree.Insert("app")
	shared.Expect(t, tree.Search("app"), true)

	tree0 := Create()
	tree0.Insert("apple")
	shared.Expect(t, tree0.Search("apple"), true)
	shared.Expect(t, tree0.Search("app"), false)
	shared.Expect(t, tree0.StartsWith("app"), true)
	tree0.Insert("app")
	shared.Expect(t, tree0.Search("app"), true)
}
