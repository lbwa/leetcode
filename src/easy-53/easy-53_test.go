package easy53

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	assert.Equal(maxSubArray([]int{5, 4, -1, 7, 8}), 23)
}
