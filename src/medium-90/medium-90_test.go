package medium90

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{[]int(nil), {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetWithDup([]int{1, 2, 2}))
}
