package medium230

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(kthSmallest(&node{Val: 3, Left: &node{Val: 1, Right: &node{Val: 2}}, Right: &node{Val: 4}}, 1), 1)
}
