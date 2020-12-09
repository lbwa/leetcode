package medium62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uniquePaths(3, 2), 3)
	assert.Equal(uniquePaths(7, 3), 28)
	assert.Equal(uniquePathsByScrollArray(3, 2), 3)
	assert.Equal(uniquePathsByScrollArray(7, 3), 28)
}
