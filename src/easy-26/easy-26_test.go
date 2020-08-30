package easy26

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	a0, a1, a2 := []int{1, 1, 2, 2, 3, 4, 4}, []int{1, 1, 1, 2, 2, 3, 3}, []int{1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4}
	a0 = a0[:removeDuplicated(a0)]
	a1 = a1[:removeDuplicated(a1)]
	a2 = a2[:removeDuplicated(a2)]
	shared.Expect(t, a0, []int{1, 2, 3, 4})
	shared.Expect(t, a1, []int{1, 2, 3})
	shared.Expect(t, a2, []int{1, 2, 3, 4})
}
