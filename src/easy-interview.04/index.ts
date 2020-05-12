export function findNumberIn2DArray(matrix: number[][], target: number) {
  /**
   * 将二维数组转换为 x-y 笛卡尔直角坐标系
   */

  let i = matrix.length
  while (i--) {
    if (target > matrix[i][0]) {
      if (findNumberInRow(matrix[i], target)) return true
    }
    if (target === matrix[i][0]) return true
  }
}

function findNumberInRow(row: number[], target: number): boolean {
  if (row.length < 1) return false

  let pivot = row[row.length >> 1]
  if (pivot > target) return findNumberInRow(row.slice(0, pivot), target)
  if (pivot < target) return findNumberInRow(row.slice(pivot + 1), target)

  return true
}
