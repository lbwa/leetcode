package hard135

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, candy([]int{1, 0, 2}))
	assert.Equal(4, candy([]int{1, 2, 2}))
}
