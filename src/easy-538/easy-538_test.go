package easy538

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, convertBST(&node{
		Val:   5,
		Left:  &node{Val: 2},
		Right: &node{Val: 13},
	}), &node{
		Val:   18,
		Left:  &node{Val: 20},
		Right: &node{Val: 13},
	})
}
