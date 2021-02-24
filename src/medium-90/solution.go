package medium90

import "sort"

func subsetWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var dfs func([]int, int)
	dfs = func(set []int, current int) {
		ans = append(ans, append([]int(nil), set...))
		for i := current; i < n; i++ {
			if i > current && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			dfs(set, i+1)
			set = set[:len(set)-1]
		}
	}
	dfs([]int{}, 0)
	return
}
