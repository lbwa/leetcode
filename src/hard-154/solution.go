package hard154

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		index := low + (high-low)/2
		pivot := nums[index]

		/*
			推荐按以下 url 中的图片综合理解
			@see https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/solution/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-de-zui--16/
			原本数组是升序排列，那么旋转后的首项在原位置大于旋转后的末项。如：
			前 ： 1 2 3 4 5 6 7
			后 ： 4 5 6 7 1 2 3
			由上可见，before[2] < before[3], after[7] < after[0]
		*/
		// 当基准项小于末项时，表示旋转点（即最小项）在 pivot+1 ~ high 之间，因为末项不会为最小项，而基准项大于末项，所以 low 指针最大可移动到 pivot+1 项，而不仅仅是 pivot 项
		if pivot > nums[high] {
			low = index + 1
			// 当基准项小于末项时，此时基准项可能为最小项，且最小项肯定不在 pivot ~ high 之间，那么移动 high 到基准项索引
		} else if pivot < nums[high] {
			high = index
		} else {
			// 当基准项等于末项时，表示 pivot 到 high 之间存在重复项，且存在旋转点（最小值）
			high--
		}
	}
	// 指针碰撞时，即为旋转点，即为最小值
	return nums[low]
}
