package easy349

import (
	"sort"
)

/*
	左右指针移动的依据是： 在 **有序** 序列中，当 a[i] > b[j] 时，移动 b 序列的 j 指针，以使得 b[j] 的值尽可能靠近 a[i] 的值，进而才会有 a[i] == b[j] 的可能性，进而得到交集
*/
func intersection(nums1 []int, nums2 []int) (ans []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0

	var isIncluded func([]int, int) bool
	isIncluded = func(nums []int, num int) bool {
		for _, val := range nums {
			if val == num {
				return true
			}
		}
		return false
	}

	for i < len1 && j < len2 {
		current1, current2 := nums1[i], nums2[j]
		if current1 == current2 {
			if !isIncluded(ans, current1) {
				ans = append(ans, current1)
			}
			i++
			j++
		} else if current1 > current2 {
			j++
		} else {
			i++
		}
	}
	return ans
}
