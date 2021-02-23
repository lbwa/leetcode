package easy136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(singleNumber([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6}), 6)
}
