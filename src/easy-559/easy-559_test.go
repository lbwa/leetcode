package easy559

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, maxDepthBfs(&node{
		Val: 1,
		Children: []*node{
			&node{Val: 3, Children: []*node{&node{Val: 5}, &node{Val: 6}}},
			&node{Val: 2},
			&node{Val: 4},
		},
	}))
	assert.Equal(3, maxDepthDfs(&node{
		Val: 1,
		Children: []*node{
			&node{Val: 3, Children: []*node{&node{Val: 5}, &node{Val: 6}}},
			&node{Val: 2},
			&node{Val: 4},
		},
	}))
}
