package medium78

/*
	思路：
		1. 考虑在 An 项中，只有两种结果，在子即中和不在子集中，那么可用二进制表示是否在子集中。那么每个子集可对应一个长度为 n(n 为源数组的长度) 的 01 序列
		2. 那么在 m 个数位中，每个数位有两种可能，那么总共有 2 的 m 次方种可能(2 * 2 * 2 ... * 2)，那么可归纳出在 n 长度的源数组种，01 序列表示的数的范围在 0 到 2 的 n 次方减一
*/
func subsets0(nums []int) (answer [][]int) {
	n := len(nums)
	// m << n 表示 m 乘以 2 的 n 次方
	// 01 序列表示的数的范围在 0 ~ 2 的 n 次方减一，那么逐个迭代得到其中一个子集
	for mask := 0; mask < (1 << n); mask++ {
		set := []int(nil) // 使用 nil 切片而非空切片
		for i, v := range nums {
			// m >> n 表示 m 除以 2 的 n 次方

			// 以下表示 mask 除以 2 的 i 次方，并与 1 做按位与操作
			// if mask>>i&1 > 0 {
			// 	set = append(set, v)
			// }

			/*
				以下表示 mask 与 2 的 i 次方执行按位与操作，实际意义在于当位操作结果为 0 时，表示二者各个数位没有同为 1 的项，反之二者存在同为 1 的项，即存在交集

				1. 1 << i 恒定返回一个二进制低位为 0 的数
				2. 基于 1 那么 mask & (1 << i) 总是返回一个 2 的某次幂或 0
			*/
			if (mask & (1 << i)) != 0 {
				set = append(set, v)
			}
		}
		answer = append(answer, append([]int(nil), set...))
	}
	return
}

func subsets1(nums []int) (answer [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(current int) {
		if current == len(nums) {
			answer = append(answer, append([]int(nil), set...))
			return
		}
		// 选择当前项，并步进到下一项
		set = append(set, nums[current])
		dfs(current + 1)
		// 不选择当前项，并步进到下一项
		set = set[:len(set)-1]
		dfs(current + 1)
	}
	dfs(0)
	return
}

func subsets2(nums []int) (ans [][]int) {
	n := len(nums)
	var dfs func([]int, int)
	dfs = func(set []int, current int) {
		ans = append(ans, append([]int(nil), set...))
		for i := current; i < n; i++ {
			set = append(set, nums[i])
			dfs(set, i+1)
			set = set[:len(set)-1]
		}
	}
	dfs([]int{}, 0)
	return
}
