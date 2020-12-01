package medium34

import (
	"sort"
)

func searchRangeON(nums []int, target int) []int {
	first, last := -1, -1
	for i, v := range nums {
		if target == v {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return []int{first, last}
}

func searchRangeOLogNByNativeMethod(nums []int, target int) []int {
	// SearchInts searches for x in a sorted slice of ints and returns the index as specified by Search. The return value is the index to insert x if x is not present (it could be len(a)). The slice must be sorted in ascending order.
	first := sort.SearchInts(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	last := sort.SearchInts(nums, target+1) - 1
	return []int{first, last}
}

func searchRangeOLogN(nums []int, target int) []int {
	var binarySearch func([]int, int, bool) int
	// lower 表示是否需要尽可能的查找低位
	binarySearch = func(nums []int, target int, lower bool) (ans int) {
		low, high := 0, len(nums)-1
		ans = len(nums)
		for low <= high {
			mid := low + (high-low)/2
			if nums[mid] > target || (lower && (nums[mid] >= target)) {
				high = mid - 1
				ans = mid
			} else {
				low = mid + 1
			}
		}
		return
	}
	first := binarySearch(nums, target, true)
	last := binarySearch(nums, target, false) - 1
	if first <= last && last < len(nums) && nums[first] == target && nums[last] == target {
		return []int{first, last}
	}
	return []int{-1, -1}
}
