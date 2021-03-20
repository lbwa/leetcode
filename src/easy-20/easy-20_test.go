package easy20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, isValid("{}"))
	assert.Equal(false, isValid("{]"))
	assert.Equal(true, isValid(""))
	assert.Equal(false, isValid("]"))
	assert.Equal(true, isValid("{[[[{()}]]]}"))
	assert.Equal(false, isValid("{[([{()}]]]}"))

	assert.Equal(true, isValid0("{}"))
	assert.Equal(false, isValid0("{]"))
	assert.Equal(true, isValid0(""))
	assert.Equal(false, isValid0("]"))
	assert.Equal(true, isValid0("{[[[{()}]]]}"))
	assert.Equal(false, isValid0("{[([{()}]]]}"))
}
