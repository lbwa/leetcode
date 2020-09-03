package medium677

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple", 3)
	shared.Expect(t, obj.Sum("ap"), 3)
}
