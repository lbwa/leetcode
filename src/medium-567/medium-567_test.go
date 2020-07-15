package medium567

import "testing"

func expect(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf(`got %t - want %t`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, CheckInclusion("hello", "ooolleoooleh"), false)
	expect(t, CheckInclusion("ab", "eidbaooo"), true)
	expect(t, CheckInclusion("ab", "eidboaoo"), false)
}
