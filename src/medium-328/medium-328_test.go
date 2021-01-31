package medium328

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	list0 := &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4}}}}
	oddEvenList((list0))
	assert.Equal(&node{Val: 1, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 4}}}}, list0)

	list1 := &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}
	oddEvenList((list1))
	assert.Equal(&node{Val: 1, Next: &node{Val: 3, Next: &node{Val: 5, Next: &node{Val: 2, Next: &node{Val: 4}}}}}, list1)
}
