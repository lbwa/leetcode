package medium103

import (
	structures "leetcode-solutions/data-structures"
	"reflect"
	"testing"
)

func expect(t *testing.T, got, want [][]int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, zigzagLevelOrder(&structures.BinaryTreeNode{
		Val:  3,
		Left: &structures.BinaryTreeNode{Val: 9},
		Right: &structures.BinaryTreeNode{
			Val:   20,
			Left:  &structures.BinaryTreeNode{Val: 15},
			Right: &structures.BinaryTreeNode{Val: 7},
		},
	}), [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	})
}
