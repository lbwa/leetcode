package medium1669

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal((*node)(nil), mergeInBetween((*node)(nil), 0, 1, &node{Val: 0}))
	assert.Equal(
		&node{Val: 0, Next: &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 1000000, Next: &node{Val: 100001, Next: &node{Val: 100002, Next: &node{Val: 5}}}}}}},
		mergeInBetween(
			&node{Val: 0, Next: &node{Val: 1, Next: &node{Val: 2, Next: &node{Val: 3, Next: &node{Val: 4, Next: &node{Val: 5}}}}}},
			3,
			4,
			&node{Val: 1000000, Next: &node{Val: 100001, Next: &node{Val: 100002}}},
		),
	)
}
