package medium62

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, uniquePaths(3, 2), 3)
	expect(t, uniquePaths(7, 3), 28)
}
