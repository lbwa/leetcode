package medium34

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(searchRangeON([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(searchRangeON([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(searchRangeON([]int{}, 0), []int{-1, -1})

	assert.Equal(searchRangeOLogNByNativeMethod([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(searchRangeOLogNByNativeMethod([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(searchRangeOLogNByNativeMethod([]int{}, 0), []int{-1, -1})

	assert.Equal(searchRangeOLogN([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(searchRangeOLogN([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(searchRangeOLogN([]int{}, 0), []int{-1, -1})

	assert.Equal(searchRangeOLogN0([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(searchRangeOLogN0([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(searchRangeOLogN0([]int{}, 0), []int{-1, -1})
}
