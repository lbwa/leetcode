package medium912

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray([]int{5, 1, 1, 2, 0, 0}))
	assert.Equal([]int{0, 1, 2, 3, 5}, sortArray([]int{5, 2, 3, 1, 0}))
}
