package easy83

import (
	structures "leetcode-solutions/data-structures"
	"reflect"
	"testing"
)

func expect(t *testing.T, got, want *structures.LinkedListNode) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %+v, want %+v`, *got, *want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, deleteDuplicated(&structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 1,
			Next: &structures.LinkedListNode{
				Val: 2,
			},
		},
	}), &structures.LinkedListNode{
		Val: 1,
		Next: &structures.LinkedListNode{
			Val: 2,
		},
	})
}
