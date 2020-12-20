package easy700

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(&node{Val: 2, Left: &node{Val: 1}, Right: &node{Val: 3}},
		searchBST(
			&node{Val: 4,
				Left:  &node{Val: 2, Left: &node{Val: 1}, Right: &node{Val: 3}},
				Right: &node{Val: 7}},
			2))
}
