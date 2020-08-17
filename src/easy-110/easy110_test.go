package easy110

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, isBalanced(&node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 3,
				Left: &node{
					Val: 4,
				},
			},
		},
		Right: &node{
			Val: 5,
		},
	}), false)

	shared.Expect(t, isBalanced(&node{
		Val: 1,
		Left: &node{
			Val: 2,
			Left: &node{
				Val: 3,
			},
			Right: &node{
				Val: 3,
				Left: &node{
					Val: 4,
				},
				Right: &node{
					Val: 4,
				},
			},
		},
		Right: &node{
			Val: 2,
		},
	}), false)

	shared.Expect(t, isBalanced(&node{
		Val: 1,
		Right: &node{
			Val: 2,
			Right: &node{
				Val: 3,
			},
		},
	}), false)

	shared.Expect(t, isBalanced(&node{
		Val: 1,
		Right: &node{
			Val: 2,
		},
	}), true)

	shared.Expect(t, isBalanced(&node{
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
			Val: 2,
			Left: &node{
				Val: 3,
			},
			Right: &node{
				Val: 4,
			},
		},
	}), true)
}
