package easy35

// SearchInsert is used to search insert position -- O(logn)
func SearchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	pivot := start + (end-start)/2

	for start < end {
		pivot = start + (end-start)/2
		current := nums[pivot]
		if current > target {
			end = pivot - 1
		} else if current < target {
			start = pivot + 1
		} else if current == target {
			return pivot
		}
	}

	if nums[start] < target {
		return start + 1
	}
	return start
}
