package medium260

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(singleNumber([]int{1, 2, 3, 2, 1, 4}), []int{4, 3})
}
