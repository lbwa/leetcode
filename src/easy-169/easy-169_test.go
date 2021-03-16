package easy169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, majorityElement([]int{3, 2, 3}))
	assert.Equal(2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
