package easy118

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, got, want [][]int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func TestSolution(t *testing.T) {
	expect(t, generate(0), [][]int{})
	expect(t, generate(1), [][]int{[]int{1}})
	/*
		[
				 [1],
				[1,1],
			 [1,2,1],
			[1,3,3,1],
		 [1,4,6,4,1]
		]
	*/
	expect(t, generate(5), [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	})
}
