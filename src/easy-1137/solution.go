package easy1137

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func tribonacci0(n int) int {
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		next := a + b + c
		a = b
		b = c
		c = next
	}
	return a
}
