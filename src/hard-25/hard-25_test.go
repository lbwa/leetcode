package hard25

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, reverseKgroup(&node{
		Val: 1,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 3,
				Next: &node{
					Val: 4,
				},
			},
		},
	}, 2), &node{
		Val: 2,
		Next: &node{
			Val: 1,
			Next: &node{
				Val: 4,
				Next: &node{
					Val: 3,
				},
			},
		},
	})

	shared.Expect(t, reverseKgroup(&node{
		Val: 1,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 3,
				Next: &node{
					Val:  4,
					Next: &node{Val: 5},
				},
			},
		},
	}, 3), &node{
		Val: 3,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 1,
				Next: &node{
					Val:  4,
					Next: &node{Val: 5},
				},
			},
		},
	})

	shared.Expect(t, reverseKgroup(&node{
		Val: 1,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 3,
			},
		},
	}, 2), &node{
		Val: 2,
		Next: &node{
			Val: 1,
			Next: &node{
				Val: 3,
			},
		},
	})

	shared.Expect(t, reverseKgroup(&node{
		Val: 1,
	}, 2), &node{
		Val: 1,
	})
}
