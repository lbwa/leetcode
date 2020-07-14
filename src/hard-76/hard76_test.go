package hard76

import (
	"testing"
)

func expect(t *testing.T, got, want string) {
	if got != want {
		t.Errorf(`got %s, want %s`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, MinWindow("ADOBECODEBANC", "ABC"), `BANC`)
	expect(t, MinWindow("aa", "aa"), `aa`)
}
