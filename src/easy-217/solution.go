package easy217

func containsDuplicate(nums []int) bool {
	items := map[int]int{}

	for _, num := range nums {
		if items[num] > 0 {
			return true
		}
		items[num]++
	}

	return false
}
