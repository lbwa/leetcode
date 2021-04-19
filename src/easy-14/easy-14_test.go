package easy14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("fl", longestCommonPrefix([]string{`flower`, `flow`, `flight`}))
	assert.Equal("", longestCommonPrefix([]string{`dog`, `racecar`, `car`}))
	assert.Equal("ab", longestCommonPrefix([]string{`abb`, `abc`}))
}
