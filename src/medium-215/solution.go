package medium215

func findKthLargest(nums []int, k int) int {
	sort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func sort(nums []int, low, high int) {
	if low >= high {
		return
	}
	index := partition(nums, low, high)
	sort(nums, low, index-1)
	sort(nums, index+1, high)
}

func partition(nums []int, low, high int) int {
	swap(nums, low+((high-low)>>1), high)
	pivot, slow := nums[high], low
	for fast := slow; fast < high; fast++ {
		if nums[fast] > pivot {
			swap(nums, fast, slow)
			slow++
		}
	}
	swap(nums, slow, high)
	return slow
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
