package medium75

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	nums0 := []int{2, 0, 2, 1, 1, 0}
	nums1 := []int{2, 0, 2, 1, 1, 0}
	nums2 := []int{2, 0, 2, 1, 1, 0}
	nums3 := []int{2, 0, 2, 1, 1, 0}
	sortColors0(nums0)
	sortColors1(nums1)
	sortColors2(nums2)
	sortColors3(nums3)
	shared.Expect(t, nums0, []int{0, 0, 1, 1, 2, 2})
	shared.Expect(t, nums1, []int{0, 0, 1, 1, 2, 2})
	shared.Expect(t, nums2, []int{0, 0, 1, 1, 2, 2})
	shared.Expect(t, nums3, []int{0, 0, 1, 1, 2, 2})
}
