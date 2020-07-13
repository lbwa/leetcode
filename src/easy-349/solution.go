package easy349

import (
	"sort"
)

// TwoPointers for resolving easy 349
func TwoPointers(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	i, j := 0, 0
	answer := []int{}

	for i < length1 && j < length2 {
		if nums1[i] == nums2[j] {
			if !has(answer, nums1[i]) {
				answer = append(answer, nums1[i])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		}
	}

	return answer
}

func has(nums []int, val int) bool {
	for _, v := range nums {
		if val == v {
			return true
		}
	}

	return false
}
