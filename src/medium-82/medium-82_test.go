package medium82

import (
	structures "leetcode-solutions/data-structures"
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	shared.Expect(t, deleteDuplicated(&structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 2,
			Next: &structures.LinkedListNode{
				Val: 3,
				Next: &structures.LinkedListNode{
					Val: 3,
					Next: &structures.LinkedListNode{
						Val: 4,
					},
				},
			},
		}}), &structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 2,
			Next: &structures.LinkedListNode{
				Val: 4,
			},
		}})
}
