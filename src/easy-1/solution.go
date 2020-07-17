package easy1

// TwoSum is used to solve two sum problem
func TwoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[target-nums[i]]; ok {
			return []int{numsMap[target-nums[i]], i}
		}
		numsMap[nums[i]] = i
	}

	return []int{}
}
