package easy349

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	got1 := TwoPointers([]int{1, 2, 2, 1}, []int{2, 2})
	want1 := []int{2}
	got2 := TwoPointers([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	want2 := []int{4, 9}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf(`%d - %d`, got1, want1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf(`%d - %d`, got2, want2)
	}
}
