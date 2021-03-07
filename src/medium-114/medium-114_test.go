package medium114

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	tree0 := &node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 3,
			},
			Right: &node{
				Val: 4,
			},
		},
		Right: &node{
			Val: 5,
			Right: &node{
				Val: 6,
			},
		},
	}
	tree1 := &node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 3,
			},
			Right: &node{
				Val: 4,
			},
		},
		Right: &node{
			Val: 5,
			Right: &node{
				Val: 6,
			},
		},
	}
	flatten0(tree0)
	flatten1(tree1)
	assert.Equal(tree0, &node{
		Val: 1,
		Right: &node{
			Val: 2,
			Right: &node{
				Val: 3,
				Right: &node{
					Val: 4,
					Right: &node{
						Val: 5,
						Right: &node{
							Val: 6,
						},
					},
				},
			},
		},
	})
	assert.Equal(tree1, &node{
		Val: 1,
		Right: &node{
			Val: 2,
			Right: &node{
				Val: 3,
				Right: &node{
					Val: 4,
					Right: &node{
						Val: 5,
						Right: &node{
							Val: 6,
						},
					},
				},
			},
		},
	})
}
