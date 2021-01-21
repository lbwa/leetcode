package easy1290

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, getDecimalValue(&node{Val: 1, Next: &node{Val: 0, Next: &node{Val: 1}}}))
	assert.Equal(5, getDecimalValue0(&node{Val: 1, Next: &node{Val: 0, Next: &node{Val: 1}}}))
}
