package medium22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	assert.Equal(generateParenthesis(1), []string{"()"})

	assert.Equal(generateParenthesis0(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	assert.Equal(generateParenthesis0(1), []string{"()"})
}
