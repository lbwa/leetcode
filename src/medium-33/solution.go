package medium33

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		// 中点在旋转数组前半部分递增序列
		if nums[low] <= nums[mid] {
			// target 值在 low ~ mid 之间
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
			// 中点在旋转数组后半部分递增序列
		} else {
			// target 在 mid ~ high 之间
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
