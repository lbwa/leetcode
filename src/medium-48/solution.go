package medium48

func rotate(matrix [][]int) {
	n := len(matrix)
	// 上下镜像翻转，即二维数组中顺序和逆序对称交换
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 将矩阵转置，即行列互换，表现为对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
