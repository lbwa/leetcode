package easy141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	node1 := &node{Val: 1}
	node2 := &node{Val: 2}
	node3 := &node{Val: 3}
	node4 := &node{Val: 4}
	assert.Equal(false, hasCycle(node1))
	node1.Next = node2
	node2.Next = node1
	assert.Equal(true, hasCycle(node1))
	node3.Next = node2
	node2.Next = node1
	node1.Next = node4
	node4.Next = node2
	assert.Equal(true, hasCycle(node1))
}
