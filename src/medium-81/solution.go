package medium81

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + ((right - left + 1) >> 1)

		if nums[mid] < nums[right] {
			if nums[mid] <= target && target <= nums[right] {
				left = mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[right] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid
			}
			// 比 medium-33 多了一个 else 分支
		} else {
			if nums[right] == target {
				return true
			}
			right--
		}
	}

	return nums[left] == target
}

func search0(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + ((right - left) >> 1)

		if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else if nums[mid] > nums[right] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
			// 比 medium-33 多了一个 else 分支
		} else {
			if nums[right] == target {
				return true
			}
			right--
		}
	}

	return nums[left] == target
}
