package easy344

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	s0 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s0)
	shared.Expect(t, s0, []byte{'o', 'l', 'l', 'e', 'h'})
}
