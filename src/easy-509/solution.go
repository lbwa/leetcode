package easy509

func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib0(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		next := a + b
		a = b
		b = next
	}
	return a
}
