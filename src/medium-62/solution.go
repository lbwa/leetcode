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

// 核心思考路径同上，不同之处在于通过滚动数组降低空间复杂度为 O(n)
func uniquePathsByScrollArray(m, n int) int {
	lastRow := make([]int, n, n) // []int{0, 0, 0, ..., 0}
	for i := range lastRow {
		lastRow[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			lastRow[j] = lastRow[j-1] + lastRow[j]
		}
	}

	return lastRow[n-1]
}
