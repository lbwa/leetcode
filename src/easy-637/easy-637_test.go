package easy637

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, averageOfLevels(&node{Val: 3, Left: &node{Val: 9}, Right: &node{Val: 20, Left: &node{Val: 15}, Right: &node{Val: 7}}}), []float64{3, 14.5, 11})
}
