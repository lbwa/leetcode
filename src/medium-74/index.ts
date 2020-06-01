export function searchMatrix(matrix: number[][], target: number): boolean {
  if (matrix.length <= 0) return false

  let row = matrix.length - 1
  let col = 0

  while (row >= 0 && col < matrix[0].length) {
    const current = matrix[row][col]
    if (current > target) {
      row--
    } else if (current < target) {
      col++
    } else if (current === target) {
      return true
    }
  }

  return false
}
