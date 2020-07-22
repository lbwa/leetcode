package hard154

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, findMin([]int{2, 2, 2, 0, 1}), 0)
	expect(t, findMin([]int{1, 3, 5}), 1)
}
