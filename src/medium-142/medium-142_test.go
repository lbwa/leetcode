package medium142

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	node0 := &node{Val: 0}
	node2 := &node{Val: 2, Next: node0}
	node3 := &node{Val: 3, Next: node2}
	node4 := &node{Val: 4, Next: node2}
	node0.Next = node4
	shared.Expect(t, detectCycle(node3), node2)
}
