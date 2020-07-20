package medium16

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, threeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
	expect(t, threeSumClosest([]int{0, 2, 1, -3}, 1), 0)
}
