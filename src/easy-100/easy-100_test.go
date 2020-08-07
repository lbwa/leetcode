package easy100

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, isSameTree(&node{
		Val:   1,
		Left:  &node{Val: 2},
		Right: &node{Val: 3},
	}, &node{
		Val:   1,
		Left:  &node{Val: 2},
		Right: &node{Val: 3},
	}), true)
}
