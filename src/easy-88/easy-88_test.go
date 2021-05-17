package easy88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)

	nums1, nums2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(nums1, []int{1, 2, 2, 3, 5, 6})
}
