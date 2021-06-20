package medium33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 5), 1)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 6), 2)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 7), 3)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	assert.Equal(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)

	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)
	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 5), 1)
	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 6), 2)
	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 7), 3)
	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	assert.Equal(search0([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
}
