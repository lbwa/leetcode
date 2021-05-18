package easy53

func maxSubArray(nums []int) int {
	prev, ans := 0, nums[0]
	var max func(a, b int) int = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, num := range nums {
		// 此处比较的是 f(i-1) 动态规划的结果，并非 0 ~ i-1 的整个序列和。因为我们根据动态规划，始终可以认为 f(i) = max(f(i-1) + nums[i], nums[i])
		// 前面累加是否有自己大，有则加进去，反之，直接从自己开始重新计和
		prev = max(prev+num, num)
		ans = max(prev, ans)
	}

	return ans
}
