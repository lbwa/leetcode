package medium47

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	set, n := []int{}, len(nums)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(current int) {
		if current == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		for i, num := range nums {
			// 1. 从左往右第一个未填入的数字
			if (i > 0 && !visited[i-1] && nums[i-1] == num) || visited[i] {
				continue
			}
			set = append(set, num)
			visited[i] = true
			dfs(current + 1)
			set = set[:len(set)-1]
			visited[i] = false
		}
	}
	dfs(0)
	return
}
