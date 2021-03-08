package medium238

func productExceptSelfOn(nums []int) []int {
	n := len(nums)
	l, r, ans := make([]int, n), make([]int, n), make([]int, n)
	l[0], r[n-1] = 1, 1

	for i := 1; i < n; i++ {
		l[i] = nums[i-1] * l[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = l[i] * r[i]
	}

	return ans
}

func productExceptSelfO1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1

	for i := 1; i < n; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}

	return ans
}
