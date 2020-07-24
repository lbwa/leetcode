package easy119

/*
	解题核心思路同 118，不同之处在于通过滚动数组实现 DP 的空间复杂度优化
*/
func getRow(rowIndex int) []int {
	// 第 i 行有 i + 1 项
	row := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		before := make([]int, rowIndex+1)
		copy(before, row)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = before[j-1] + before[j]
		}
	}
	return row
}
