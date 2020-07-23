package medium64

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}), 7)
	expect(t, minPathSum([][]int{
		[]int{5},
	}), 5)
	expect(t, minPathSum([][]int{}), 0)
	expect(t, minPathSum([][]int{[]int{}}), 0)
}
