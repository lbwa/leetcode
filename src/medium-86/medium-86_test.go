package medium86

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(&node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 2, Next: &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 5}}}}}}, partition(&node{Val: 1, Next: &node{Val: 4, Next: &node{Val: 3, Next: &node{Val: 2, Next: &node{Val: 5, Next: &node{Val: 2}}}}}}, 3))
}
