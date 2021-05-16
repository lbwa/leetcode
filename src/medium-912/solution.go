package medium912

func sortArray(nums []int) []int {
	var swap func([]int, int, int) = func(nums []int, a, b int) {
		nums[a], nums[b] = nums[b], nums[a]
	}

	var partition func([]int, int, int) int = func(nums []int, start, end int) int {
		pivotIndex := start + ((end - start) >> 1)
		pivot := nums[pivotIndex]
		swap(nums, pivotIndex, end)
		lowIndex := start
		for i := start; i < end; i++ {
			if nums[i] < pivot {
				swap(nums, i, lowIndex)
				lowIndex++
			}
		}
		swap(nums, lowIndex, end)
		return lowIndex
	}

	var sort func([]int, int, int)
	sort = func(nums []int, start, end int) {
		if start >= end {
			return
		}
		index := partition(nums, start, end)
		sort(nums, start, index-1)
		sort(nums, index+1, end)
	}

	sort(nums, 0, len(nums)-1)
	return nums
}
