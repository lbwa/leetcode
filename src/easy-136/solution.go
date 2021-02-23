package easy136

func singleNumber(nums []int) int {
	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
