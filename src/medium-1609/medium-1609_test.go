package medium1609

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, isEvenOddTree(&node{Val: 1}), true)
	shared.Expect(t, isEvenOddTree(&node{
		Val: 5,
		Left: &node{
			Val:   9,
			Left:  &node{Val: 3},
			Right: &node{Val: 5},
		},
		Right: &node{
			Val:  1,
			Left: &node{Val: 7},
		},
	}), false)
}
