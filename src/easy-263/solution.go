package easy263

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	factors := []int{2, 3, 5}
	for _, factor := range factors {
		for n%factor == 0 {
			n /= factor
		}
	}

	return n == 1
}

// 法二；最小堆
// 法三：动态规划
