package easy367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, firstUniqChar("leetcode"))
	assert.Equal(2, firstUniqChar("loveleetcode"))

	assert.Equal(0, firstUniqCharFromSlice("leetcode"))
	assert.Equal(2, firstUniqCharFromSlice("loveleetcode"))
}
