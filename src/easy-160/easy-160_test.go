package easy160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	list0 := &node{Val: 4, Next: &node{Val: 1, Next: &node{Val: 8, Next: &node{Val: 4, Next: &node{Val: 5}}}}}
	list1 := &node{Val: 5, Next: &node{Val: 0, Next: &node{Val: 1, Next: list0.Next.Next}}}
	assert.Equal(list0.Next.Next, getIntersectionNode(list0, list1))
	assert.Equal(list0.Next.Next, getIntersectionNode(list1, list0))
}
