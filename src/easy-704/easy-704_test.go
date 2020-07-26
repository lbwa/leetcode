package easy704

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, search([]int{-1, 0, 3, 5, 9, 12}, 13), -1)
	expect(t, search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
}
