package medium406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})

	var ans [][]int
	for _, person := range people {
		i := person[1]
		ans = append(ans[:i], append([][]int{person}, ans[i:]...)...)
	}
	return ans
}
