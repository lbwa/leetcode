package medium143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	list0 := &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4}}}}
	reorderList(list0)
	assert.Equal(&node{Val: 1, Next: &node{Val: 4, Next: &node{Val: 2, Next: &node{Val: 3}}}}, list0)

	list1 := &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}
	reorderList(list1)
	assert.Equal(&node{Val: 1, Next: &node{Val: 5, Next: &node{Val: 2, Next: &node{Val: 4, Next: &node{Val: 3}}}}}, list1)
}
