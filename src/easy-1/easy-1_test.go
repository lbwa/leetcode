package easy1

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, TwoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	expect(t, TwoSum([]int{3, 2, 4}, 6), []int{1, 2})
}
