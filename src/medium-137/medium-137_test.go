package medium137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(singleNumber([]int{2, 2, 3, 2}), 3)
	assert.Equal(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}), 99)
}
