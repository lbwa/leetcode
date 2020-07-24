package easy119

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
	expect(t, getRow(0), []int{1})
	expect(t, getRow(1), []int{1, 1})
	expect(t, getRow(2), []int{1, 2, 1})
	expect(t, getRow(3), []int{1, 3, 3, 1})
	expect(t, getRow(4), []int{1, 4, 6, 4, 1})
}
