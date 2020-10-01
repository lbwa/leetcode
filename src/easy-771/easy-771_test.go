package easy771

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, numJewelsInStones("aA", "aAAbbbb"), 3)
	shared.Expect(t, numJewelsInStones("z", "ZZ"), 0)
}
