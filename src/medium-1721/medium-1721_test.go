package medium1721

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(&node{Val: 1, Next: &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 5}}}}}, swapNodes(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}, 2))

	assert.Equal((*node)(nil), swapNodes((*node)(nil), 0))

	assert.Equal(&node{Val: 1, Next: &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 5}}}}}, swapNodesVal(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}, 2))

	assert.Equal((*node)(nil), swapNodesVal((*node)(nil), 0))
}
