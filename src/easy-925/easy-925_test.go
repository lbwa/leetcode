package easy925

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, isLongPressedName("alex", "aaleex"), true)
	assert.Equal(t, isLongPressedName("saeed", "ssaaedd"), false)
	assert.Equal(t, isLongPressedName("leelee", "lleeelee"), true)
	assert.Equal(t, isLongPressedName("laiden", "laiden"), true)
}
