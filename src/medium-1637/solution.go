package medium1637

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(a, b int) bool {
		return points[a][0] < points[b][0]
	})
	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var ans int
	for i := 0; i < len(points)-1; i++ {
		current, after := points[i], points[i+1]
		ans = max(ans, after[0]-current[0])
	}
	return ans
}
