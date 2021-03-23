package medium17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal([]string{"a","b","c"}, letterCombinations("2"))
	assert.Equal([]string{}, letterCombinations(""))
}
