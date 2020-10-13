package easy530

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, getMinimumDifference(&node{
		Val:   1,
		Right: &node{Val: 3, Left: &node{Val: 2}},
	}), 1)
}
