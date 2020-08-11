package medium11

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, maxArea([]int{2, 3, 4, 5, 18, 17, 6}), 17)
	shared.Expect(t, maxArea([]int{}), 0)
	shared.Expect(t, maxArea([]int{1}), 0)
	shared.Expect(t, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}
