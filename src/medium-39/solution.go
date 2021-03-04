package medium39

func combinationSum(candidates []int, target int) (answer [][]int) {
	comb := []int{}
	var dfs func(int, int)
	dfs = func(target, index int) {
		// break up dfs recursion when we reach tail node
		if index == len(candidates) {
			return
		}
		// break up dfs recursion and save result when we reach target sum
		if target == 0 {
			answer = append(answer, append([]int(nil), comb...))
			return
		}
		/*
			We have all 2 decisions, use current element or not.
		*/
		// make decision 1, skip `index` elements
		dfs(target, index+1)
		// make decision 2, use `index`th element
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			// remove element when we have made decision(because we use the only one variable to store all possible results)
			comb = comb[:len(comb)-1]
		}
	}
	// activate dfs to find all results
	dfs(target, 0)
	return
}

func combinationSum1(candidates []int, target int) (answer [][]int) {
	var dfs func(int, int, *[]int)
	dfs = func(target, index int, combine *[]int) {
		if index >= len(candidates) {
			return
		}
		if target == 0 {
			answer = append(answer, append([]int(nil), *combine...))
			return
		}
		dfs(target, index+1, combine)
		if target-candidates[index] >= 0 {
			// we should use a new slice to store next result when we are using current node.
			comb := append(*combine, candidates[index])
			dfs(target-candidates[index], index, &comb)
		}
	}
	dfs(target, 0, &[]int{})
	return
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	set := []int{}
	n := len(candidates)
	var sum int
	var dfs func(int)
	dfs = func(current int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		for i := current; i < n; i++ {
			num := candidates[i]
			set = append(set, num)
			sum += num
			dfs(i)
			set = set[:len(set)-1]
			sum -= num
		}
	}
	dfs(0)
	return
}
