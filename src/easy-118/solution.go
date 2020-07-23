package easy118

func generate(numRows int) [][]int {
	/*
		1. 第 i 行有 i + 1 项
		2. dp[i][j] = dp[i-1][j-1]+dp[i-1][j]
	*/
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][(i+1)-1] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}
