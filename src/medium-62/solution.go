package medium62

func uniquePaths(m, n int) int {
	rows, cols, matrix := n, m, make([][]int, n)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		matrix[i][0] = 1
	}
	for i := 0; i < cols; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[n-1][m-1]
}
