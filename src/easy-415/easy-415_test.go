package easy415

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, addStrings(`0`, `0`), `0`)
	shared.Expect(t, addStrings(`123`, `1234`), `1357`)
	shared.Expect(t, addStrings(`12345`, `1234`), `13579`)
}
