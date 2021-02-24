package medium78

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{
		[]int(nil), // nil 切片与空切片不等价
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}, subsets0([]int{1, 2, 3}))
	assert.Equal([][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, []int(nil)}, subsets1([]int{1, 2, 3}))
	assert.Equal([][]int{[]int(nil), {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}, subsets2([]int{1, 2, 3}))
}
