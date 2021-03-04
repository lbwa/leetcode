package medium39

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{7}, {2, 2, 3}})
	shared.Expect(t, combinationSum1([]int{2, 3, 6, 7}, 7), [][]int{{7}, {2, 2, 3}})
	shared.Expect(t, combinationSum2([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}})
}
