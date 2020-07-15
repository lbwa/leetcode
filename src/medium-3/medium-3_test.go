package medium3

import "testing"

func expect(t *testing.T, got, want int) {
	if got != want {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, LengthOfLongestSubstring("abcabcbb"), 3)
	expect(t, LengthOfLongestSubstring("bbbbb"), 1)
	expect(t, LengthOfLongestSubstring("pwwkew"), 3)
}
