package easy206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	head := &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4}}}}
	want := &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 1}}}}
	assert.Equal(want, reverseListIteration(head))
	assert.Equal(want, reverseListRecursion(head))
}
