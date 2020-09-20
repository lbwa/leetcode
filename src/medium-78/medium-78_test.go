package medium78

import (
	"leetcode-solutions/src/shared"
	"reflect"
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
	got1 := subsets1([]int{1, 2, 3})
	want1 := [][]int{
		[]int(nil), // nil 切片与空切片不等价
		[]int{1},
		[]int{2},
		[]int{1, 2},
		[]int{3},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	}
	for _, got := range got1 {
		if !isIncluded(want1, got) {
			t.Errorf(`%d should be in answer`, want1)
			return
		}
	}
}

func isIncluded(nums [][]int, val []int) bool {
	for _, v := range nums {
		if reflect.DeepEqual(v, val) {
			return true
		}
	}
	return false
}
