package medium46

func permute(nums []int) (ans [][]int) {
	var isInSlice func([]int, int) bool
	isInSlice = func(list []int, num int) bool {
		for _, item := range list {
			if item == num {
				return true
			}
		}
		return false
	}

	set, n := []int{}, len(nums)
	var dfs func(int)
	dfs = func(current int) {
		// 在 backtrack 模板中，首先定义递归退出条件，退出时要将路径添加到结果中
		if current == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		for _, num := range nums {
			// 原始序列不含重复值且求全排列，故无重复值添加
			if isInSlice(set, num) {
				continue
			}
			// 选择当前项
			set = append(set, num)
			// 在选择当前项后，进入下一决策
			dfs(current + 1)
			// 回溯当前项
			set = set[:len(set)-1]
		}
	}
	dfs(0)
	return
}
