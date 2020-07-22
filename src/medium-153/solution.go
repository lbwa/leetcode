package medium153

/* 另一种二分查找实现见 hard-154（推荐查看 hard-154 解法更为简洁，且支持存在重复值的查找） */

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	/*
		原本数组是升序排列，那么旋转后的首项在原位置大于旋转后的末项。如：
		前 ： 1 2 3 4 5 6 7
		后 ： 4 5 6 7 1 2 3
		由上可见，before[2] < before[3], after[7] < after[0]
	*/
	/*
		当旋转后的数组 nums[0] < nums[right] 时，表示此时数组已经排序，故 nums[0] 为最小值
	*/
	if nums[0] < nums[right] {
		return nums[0]
	}

	for left <= right {
		pivot := left + (right-left)/2

		/*
			根据原数组是升序的特点，升序排列应该是前项始终不大于后项，那么当前项大于后项时，那么后项即为旋转点，那么此时的后项为原升序数组中的最小值
		*/
		// 注意必须搜先判断 pivot 与 后项是否相等，之后再判断与 pivot-1 项，因为 pivot 可为 0
		if nums[pivot] > nums[pivot+1] {
			return nums[pivot+1]
		}
		// 同上 if 语句
		if nums[pivot-1] > nums[pivot] {
			return nums[pivot]
		}
		/*
			基于原数组的升序排列，且基于前两个 pivot 判断，pivot 不为断开项时，
			那么当 pivot 项大于首项时，那么至少在 0 ~ pivot 项之间是不存在断点项的，那么右移 left 至 pivot + 1 项
		*/
		if nums[pivot] > nums[0] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return -1
}
