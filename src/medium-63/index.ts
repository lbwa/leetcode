export function uniquePathsWithObstaclesSpaceO2n(
  obstacleGrid: number[][]
): number {
  const rows = obstacleGrid.length
  const cols = obstacleGrid[0].length
  const dp: number[][] = []
  for (let i = 0; i < rows; i++) dp[i] = []
  // 此处有个细节是，若在迭代中存在 1 值（即障碍物），那么将打断之后的赋值，因为之后的路径无法到达终点
  for (let i = 0; i < rows && obstacleGrid[i][0] === 0; i++) dp[i][0] = 1
  for (let i = 0; i < cols && obstacleGrid[0][i] === 0; i++) dp[0][i] = 1

  // 根据动态方程得到以下迭代
  for (let i = 1; i < rows; i++) {
    for (let j = 1; j < cols; j++) {
      if (obstacleGrid[i][j] === 0) {
        dp[i][j] = (dp[i - 1][j] || 0) + (dp[i][j - 1] || 0)
      }
    }
  }

  return dp[rows - 1][cols - 1] || 0
}

/**
 * 借助 `滚动数组` 优化空间复杂度
 */
export function uniquePathsWithObstaclesSpaceOn(
  obstacleGrid: number[][]
): number {
  const rows = obstacleGrid.length
  const cols = obstacleGrid[0].length
  const currentRow: number[] = new Array(cols).fill(0)

  // 确保起点不是障碍物
  currentRow[0] = obstacleGrid[0][0] === 0 ? 1 : 0

  // 根据动态方程得到以下循环
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (obstacleGrid[i][j] !== 1) {
        if (j - 1 >= 0 && obstacleGrid[i][j - 1] === 0) {
          // 此处缩写一定要保证 currentRow 的相加相有值，否则使用备选值，如
          // currentRow[j] = (currentRow[j] || 0) + (currentRow[j - 1] || 0)
          currentRow[j] += currentRow[j - 1]
        }
      } else {
        currentRow[j] = 0
      }
    }
  }

  return currentRow[cols - 1]
}
