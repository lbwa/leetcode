/**
 * 完全基于动态方程的逻辑细节的解法
 * @时间复杂度 O(m * n) // 因为需要经过每个节点，故为 m * n
 * @空间复杂度 O(m * n) // 因为需要建立额外的结果矩阵存储结果，故为 m * n
 * @动态方程 在 (m, n) 点时，需要的步数为 (m - 1, n) 点 和 (m, n - 1) 点之和
 * - if (col && row), f(col, row) = f(col - 1, row) + f(col, row - 1)
 * - if (col === 0), f(0, row) = 1
 * - if (row === 0), f(col, 0) = 1
 */
export function uniquePathsSpaceO2n(col: number, row: number): number {
  const dp: number[][] = []
  for (let i = 0; i < row; i++) dp[i] = []

  for (let i = 0; i < col; i++) dp[0][i] = 1
  for (let i = 0; i < row; i++) dp[i][0] = 1
  for (let i = 1; i < row; i++) {
    for (let j = 1; j < col; j++) {
      dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
    }
  }
  return dp[row - 1][col - 1]
}

export function uniquePathsSpaceOn(col: number, row: number) {
  const currentRow = new Array<number>(col).fill(1)

  for (let i = 1; i < row; i++) {
    for (let j = 1; j < col; j++) {
      currentRow[j] += currentRow[j - 1]
    }
  }

  return currentRow[col - 1]
}
