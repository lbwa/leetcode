package easy191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, hammingWeight(0b1011))
	assert.Equal(6, hammingWeight(0b10111011))
}
