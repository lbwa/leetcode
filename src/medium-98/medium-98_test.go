package medium98

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, isValidBSTByIteration(nil), true)
	shared.Expect(t, isValidBSTByIteration(&node{Val: 2, Left: &node{Val: 1}, Right: &node{Val: 3}}), true)
	shared.Expect(t, isValidBSTByIteration(&node{Val: 1, Left: &node{Val: 1}}), false)

	shared.Expect(t, isValidBSTByRecursion(nil), true)
	shared.Expect(t, isValidBSTByRecursion(&node{Val: 2, Left: &node{Val: 1}, Right: &node{Val: 3}}), true)
	shared.Expect(t, isValidBSTByRecursion(&node{Val: 1, Left: &node{Val: 1}}), false)
}
