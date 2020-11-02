package easy349

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
