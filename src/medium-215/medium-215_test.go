package medium215

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	shared.Expect(t, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
}
