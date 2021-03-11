package medium64

func minPathSum(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}

	rows, cols, dp := len(grid), len(grid[0]), make([][]int, len(grid))

	dp[0] = make([]int, cols)
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i] = make([]int, cols)
		// 因为每次只能向下和向右，故有以下赋值
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < cols; i++ {
		// 因为每次只能向下和向右，故有以下赋值
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][cols-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSumOptimize(grid [][]int) int {
	if grid == nil || len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}

	var min func(int, int) int
	min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp := make([]int, len(grid[0]))
	rows, cols := len(grid), len(grid[0])
	dp[0] = grid[0][0]

	for i := 1; i < cols; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < rows; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < cols; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[cols-1]
}
