package medium33

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 5), 1)
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 6), 2)
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 7), 3)
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	expect(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
}
