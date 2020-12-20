package easy206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	want := &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 1}}}}
	assert.Equal(want, reverseListIteration(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4}}}}))
	assert.Equal(want, reverseListRecursion(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4}}}}))
}
