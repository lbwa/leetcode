package easy657

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, judgeCircle("LRLRLR"), true)
	shared.Expect(t, judgeCircle("UDUDUDDLRD"), false)
}
