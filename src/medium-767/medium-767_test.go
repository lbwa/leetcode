package medium767

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("aba", reorganizeString("aab"))
	assert.Equal("vlvov", reorganizeString("vvvlo"))
}
