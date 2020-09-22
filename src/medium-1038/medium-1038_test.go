package medium1038

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, bstToGst(&node{
		Val:   5,
		Left:  &node{Val: 2},
		Right: &node{Val: 13},
	}), &node{
		Val:   5,
		Left:  &node{Val: 2},
		Right: &node{Val: 13},
	})
}
