package easy203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal((*node)(nil), removeElement(&node{Val: 1}, 1))
	assert.Nil(removeElement(&node{Val: 1}, 1))
	assert.Equal(&node{Val: 1}, removeElement(&node{Val: 1, Next: &node{Val: 2}}, 2))
}
