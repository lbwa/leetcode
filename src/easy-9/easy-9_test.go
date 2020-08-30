package easy9

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, isPalindrome(1234321), true)
	shared.Expect(t, isPalindrome(-1234321), false)
	shared.Expect(t, isPalindrome(12343212), false)
}
