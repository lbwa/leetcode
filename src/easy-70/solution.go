package easy70

func climbStairs(n int) int {
	prevPrev, prev, current := 0, 0, 1
	for i := 0; i < n; i++ {
		prevPrev = prev
		prev = current
		current = prevPrev + prev
	}
	return current
}
