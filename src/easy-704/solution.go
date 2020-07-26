package easy704

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		current := nums[mid]

		if current > target {
			high = mid - 1
		} else if current < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
