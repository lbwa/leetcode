package easy35

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, SearchInsert([]int{1, 3, 5, 6}, 5), 2)
	expect(t, SearchInsert([]int{1, 3, 5, 6}, 2), 1)
	expect(t, SearchInsert([]int{1, 3, 5, 6}, 7), 4)
	expect(t, SearchInsert([]int{1, 3, 5, 6}, 0), 0)
}
