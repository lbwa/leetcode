package medium33

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// 参考 https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/er-fen-fa-python-dai-ma-java-dai-ma-by-liweiwei141/
	for left < right {
		// 与题无关，当有边界设置为 left = mid 和 right = mid-1 出现时，需要将 mid 的下
		// 取整调整为上取整，以避免出现死循环，故有 left + ((right-left+1) >> 1)
		// 原因在于如果我们不+1，则是取下整，得到的中点 mid 即为 left，而在
		// [left,right] 区间中只有两个元素时，若再次进入 if/else 分支，由于是闭区间，又会
		// 取 mid 的值并赋给 left，从而使二者变为定值，便进入死循环。
		mid := left + ((right - left + 1) >> 1)
		pivot := nums[mid]

		// step1: 首先判断 mid 所在的单调区间，利用单调性进行二分操作
		//        数组最大值
		//           /|
		//          / |
		//         /  |
		// 数组拐点--- | ---数组拐点
		//            |   /
		//            |  /
		//            | /
		//            |/
		//       数组中最小值
		// mid 位于最小值到拐点区间，此时 [mid, right] 一定是有序的
		if pivot < nums[right] {
			if pivot <= target && target <= nums[right] {
				left = mid // 下一轮区间 [mid, right]
			} else {
				right = mid - 1 // 下一轮区间 [left, mid-1]
			}
		} else {
			// 结合 !(pivot < nums[right]) 条件，mid 此时位于拐点到最大值区间，此时 [left, mid] 一定是有序的。
			if nums[left] <= target && target <= nums[mid-1] {
				right = mid - 1 // 下一轮区间 [left, mid-1]
			} else {
				left = mid // 下一轮区间 [mid, right]
			}

			// 补充说明：由于中间数上取整，在区间只剩下两个元素的时候，mid 与 right 重合，
			// 方法逻辑走到 else 分支里。此时恰好 if 这个分支看到的是 left 和 mid - 1 ，用
			// 到的都是 == 号，等价于判断 nums[left] == target。因此依然可以缩减区间，注意
			// 这里 if 里面的 nums[left] <= target && target <= nums[mid - 1] ，
			// 不可以写成 nums[left] <= target && target < nums[mid]，如在 [3,1] 这样的
			// 用例中将出现错误答案

		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func search0(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + ((right - left) >> 1)

		// mid 位于最小值到拐点区间，此时 [mid, right] 一定是有序的
		if nums[mid] < nums[right] {
			// 此 if 中保证始终可能有 target 等于 right 边界的比较
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			// 结合 !(pivot < nums[right]) 条件，mid 此时位于拐点到最大值区间，此时 [left, mid] 一定是有序的。
			// 此 if 中保证始终可能有 target 等于 left 边界的比较
			// 此处有 target <= nums[mid] 判断是因为之前的 if 判断没有覆盖到等于 nums[mid] 的场景
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
