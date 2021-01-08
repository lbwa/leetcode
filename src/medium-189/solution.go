package medium189

func rotate(nums []int, k int) {
	extra := make([]int, len(nums))
	for i, v := range nums {
		extra[(i+k)%len(nums)] = v
	}
	copy(nums, extra)
}
