package easy242

import "testing"

func expect(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf(`got %t, want %t`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, IsAnagram(`aa`, `aacc`), false)
	expect(t, IsAnagram(`ccac`, `a`), false)
	expect(t, IsAnagram("anagram", "nagaram"), true)
	expect(t, IsAnagram("rat", "car"), false)
	expect(t, IsAnagram("a", "ab"), false)
}
