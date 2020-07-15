package medium438

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
	expect(t, FindAnagram("cbaebabacd", "abc"), []int{0, 6})
	expect(t, FindAnagram("abab", "ab"), []int{0, 1, 2})
}
