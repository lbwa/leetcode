package medium78

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, subsets0([]int{1, 2, 3}), [][]int{
		[]int(nil), // nil 切片与空切片不等价
		[]int{1},
		[]int{2},
		[]int{1, 2},
		[]int{3},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	})
}
