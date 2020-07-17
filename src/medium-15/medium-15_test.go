package medium15

import "testing"

func expect(t *testing.T, answers [][]int) {
	for _, answer := range answers {
		sum := 0
		for _, num := range answer {
			sum += num
		}
		if sum != 0 {
			t.Errorf(`got %d, sum is not zero`, answer)
		}
	}
}

func TestSolution(t *testing.T) {
	expect(t, ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}
