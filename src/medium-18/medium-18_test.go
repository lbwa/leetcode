package medium18

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	got, want := FourSum([]int{1, 0, -1, 0, -2, 2}, 0), [][]int{
		[]int{1, 2, -1, -2},
		[]int{0, 2, 0, -2},
		[]int{0, 1, 0, -1},
	}
	for i := 0; i < len(got); i++ {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Errorf(`got %d, want %d`, got[i], want[i])
		}
	}
}
