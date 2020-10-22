package easy234

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, isPalindrome(nil), true)
	assert.Equal(t, isPalindrome(&node{Val: 1, Next: &node{Val: 2}}), false)
	assert.Equal(t, isPalindrome(&node{
		Val: 1,
		Next: &node{
			Val: 2,
			Next: &node{
				Val: 2,
				Next: &node{
					Val: 1,
				},
			},
		},
	}), true)
}
