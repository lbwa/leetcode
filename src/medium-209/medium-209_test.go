package medium209

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, minSubArrayLen(100, []int{}), 0)
	expect(t, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
}
