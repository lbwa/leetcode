package medium43

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, multiply(`9`, `9`), `81`)
	shared.Expect(t, multiply(`123`, `3213214321`), `395225361483`)
}
