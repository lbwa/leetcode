package hard135

func candy(ratings []int) (ans int) {
	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(ratings)
	fromStart := make([]int, n)
	for i, rate := range ratings {
		if i > 0 && rate > ratings[i-1] {
			fromStart[i] = fromStart[i-1] + 1
		} else {
			fromStart[i] = 1
		}
	}

	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(fromStart[i], right)
	}

	return
}
