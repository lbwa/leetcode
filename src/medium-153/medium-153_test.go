package medium153

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, findMin([]int{1, 2}), 1)
	expect(t, findMin([]int{2, 1}), 1)
	expect(t, findMin([]int{3, 4, 5, 1, 2}), 1)
	expect(t, findMin([]int{4, 5, 6, 7, 0, 1, 2}), 0)
}
