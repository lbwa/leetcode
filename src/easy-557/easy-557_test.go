package easy557

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
	assert.Equal("s'teL ekat edoCteeL tsetnoc", reverseWords1("Let's take LeetCode contest"))
}
