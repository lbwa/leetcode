package medium92

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		&node{Val: 1, Next: &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 5}}}}},
		reverseBetween(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}, 2, 4),
	)
	assert.Equal(
		&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}},
		reverseBetween(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}, 0, 0),
	)
	assert.Equal(
		(*node)(nil),
		reverseBetween((*node)(nil), 0, 0),
	)
}
