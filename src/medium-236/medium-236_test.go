package medium236

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	node0 := &node{Val: 1}
	node1 := &node{Val: 2}
	node2 := &node{
		Val:   3,
		Left:  node0,
		Right: node1,
	}
	shared.Expect(t, lowestCommonAncestor(node2, node0, node1), node2)
}
