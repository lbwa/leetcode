package medium96

func numTrees(n int) int {
	num := make([]int, n+1)
	num[0], num[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			num[i] += num[j-1] * num[i-j]
		}
	}
	return num[n]
}
