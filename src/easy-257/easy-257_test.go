package easy257

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, binaryTreePaths(&node{
		Val:   1,
		Left:  &node{Val: 2, Right: &node{Val: 5}},
		Right: &node{Val: 3},
	}), []string{"1->2->5", "1->3"})
}
