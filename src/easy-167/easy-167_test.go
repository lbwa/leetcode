package easy167

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
	expect(t, twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	expect(t, twoSum([]int{5, 25, 75}, 100), []int{2, 3})
	expect(t, twoSumWithBinary([]int{2, 7, 11, 15}, 9), []int{1, 2})
	expect(t, twoSumWithBinary([]int{5, 25, 75}, 100), []int{2, 3})
}
