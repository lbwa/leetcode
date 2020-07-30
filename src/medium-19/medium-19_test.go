package medium19

import (
	structures "leetcode-solutions/data-structures"
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, removeNthFromEnd(&structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 2,
			Next: &structures.LinkedListNode{
				Val: 3,
				Next: &structures.LinkedListNode{
					Val: 4,
					Next: &structures.LinkedListNode{
						Val: 5,
					},
				},
			},
		}}, 2), &structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 2,
			Next: &structures.LinkedListNode{
				Val: 3,
				Next: &structures.LinkedListNode{
					Val: 5,
				},
			},
		}})

	if removeNthFromEnd(&structures.LinkedListNode{
		Val: 1,
	}, 1) != nil {
		t.Error(`Assert failed`)
	}
}
