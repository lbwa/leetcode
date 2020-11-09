package medium973

import "sort"

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		// 求解平方根时需要引入 float 类型，故存在误差，但仅求解平方和，即为 int 类型，故不用考虑误差
		spaceA := points[a][0]*points[a][0] + points[a][1]*points[a][1]
		spaceB := points[b][0]*points[b][0] + points[b][1]*points[b][1]
		return spaceA < spaceB
	})

	return points[:K]
}
