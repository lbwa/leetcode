package easy35

// time complexity O(logN), space complexity O(1)
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if nums[right] < target {
		return right + 1
	}

	for left < right {
		pivotIndex := left + ((right - left) >> 1)
		pivot := nums[pivotIndex]

		// 始终要保证所有的 if 语句无交集，且并集一定是 [left, right] 区间
		if pivot < target {
			left = pivotIndex + 1 // [pivotIndex + 1, right]
		} else {
			right = pivotIndex // [left, pivotIndex]
		}
	}

	return left
}
