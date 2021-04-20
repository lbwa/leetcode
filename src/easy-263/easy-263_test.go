package easy263

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.True(isUgly(1))
	assert.True(isUgly(6))
	assert.True(isUgly(8))
	assert.False(isUgly(14))
}
