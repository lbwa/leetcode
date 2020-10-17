package medium46

func permute(nums []int) [][]int {
	// 全局共享一个 track 变量用于追踪路径
	answer, track := [][]int{}, []int{}
	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 在 backtrack 模板中，首先定义递归退出条件，退出时要将路径添加到结果中
		if len(track) == len(nums) {
			path := make([]int, len(track))
			copy(path, track)
			answer = append(answer, path)
			return
		}

		for _, num := range nums {
			// 原始序列不含重复值且求全排列，故无重复值添加
			if !isInSlice(track, num) {
				// 选择当前项
				track = append(track, num)
				// 在选择当前项后，进入下一决策
				backtrack(nums)
				// 回溯当前项
				track = track[:len(track)-1]
			}
		}
	}
	backtrack(nums)
	return answer
}

func isInSlice(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}
