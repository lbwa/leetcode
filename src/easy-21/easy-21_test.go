package easy21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(&list{Val: 1, Next: &list{Val: 1, Next: &list{Val: 2, Next: &list{Val: 3, Next: &list{Val: 4, Next: &list{Val: 4}}}}}}, mergeTwoList(&list{Val: 1, Next: &list{Val: 2, Next: &list{Val: 4}}}, &list{Val: 1, Next: &list{Val: 3, Next: &list{Val: 4}}}))
}
