export function maxWidthOfVerticalArea(points: number[][]): number {
  points.sort(([a], [b]) => a - b)
  let ans = 0
  for (let i = 0; i < points.length - 1; i++) {
    const [current] = points[i]
    const [after] = points[i + 1]
    ans = Math.max(ans, after - current)
  }
  return ans
}
