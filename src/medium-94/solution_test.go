package solution

import (
	structures "leetcode-solutions/data-structures"
	"reflect"
	"testing"
)

func TestIteration(t *testing.T) {
	want := []int{1, 3, 2}
	tree := structures.BinaryTreeNode{
		Val: 1,
		Right: &structures.BinaryTreeNode{
			Val: 2,
			Left: &structures.BinaryTreeNode{
				Val: 3,
			},
		},
	}
	if got := IterativeInOrderTraversal(&tree); !reflect.DeepEqual(got, want) {
		t.Errorf(`got %q, want %q`, got, want)
	}
}

func TestRecursion(t *testing.T) {
	want := []int{1, 3, 2}
	tree := structures.BinaryTreeNode{
		Val: 1,
		Right: &structures.BinaryTreeNode{
			Val: 2,
			Left: &structures.BinaryTreeNode{
				Val: 3,
			},
		},
	}
	if got := RecursiveInOrderTraversal(&tree); !reflect.DeepEqual(got, want) {
		t.Errorf(`got %q, want %q`, got, want)
	}
}
