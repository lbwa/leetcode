package easy941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(false, validMuntainArray([]int{2, 1}))
	assert.Equal(false, validMuntainArray([]int{3, 5, 5}))
	assert.Equal(true, validMuntainArray([]int{0, 3, 2, 1}))
}
