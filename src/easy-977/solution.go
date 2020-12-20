package easy977

func sortedSqares(a []int) []int {
	left, right := 0, len(a)-1
	ans := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		low, high := a[left]*a[left], a[right]*a[right]
		if low > high {
			ans[i] = low
			left++
		} else {
			ans[i] = high
			right--
		}
	}
	return ans
}
