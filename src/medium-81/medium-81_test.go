package medium81

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.True(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.False(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	assert.True(search([]int{1, 0, 1, 1, 1}, 0))

	assert.True(search0([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.False(search0([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	assert.True(search0([]int{1, 0, 1, 1, 1}, 0))
}
