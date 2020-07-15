package medium438

/*
	推荐解决 easy 242 有效的字母异位词，来首先解决如何判断存在异位词的问题
*/

// FindAnagram is used to find all anagrams
func FindAnagram(s, p string) []int {
	needsMap, windowMap := map[byte]int{}, map[byte]int{}
	left, right := 0, 0
	answer, validsInWindow := []int{}, 0

	// 构建目标异位词的哈希映射，最大 size 为包含所有 26 个字母，即 26
	for i := range p {
		needsMap[p[i]]++
	}

	for right < len(s) {
		// 增大窗口
		current := s[right]
		right++
		// 若当前为异位词所包含的字母，那么进入处理
		if _, ok := needsMap[current]; ok {
			windowMap[current]++
			// 当指定字母频率相等时，计数一次，当 validsInWindow 与异位词的映射大小相同时，
			// 表示当前窗口存在 p 的异位词
			if needsMap[current] == windowMap[current] {
				validsInWindow++
			}
		}

		// 当窗口大于 p 时，开始尝试缩小窗口，并找到异位词
		for (right - left) >= len(p) {
			// 如 easy 242 思路，那么在以下计数相等的情况下，即在窗口中存在异位词
			if validsInWindow == len(needsMap) {
				answer = append(answer, left)
			}

			// 缩小窗口
			deleted := s[left]
			left++
			// 剔除因缩小窗口降低的字母频率
			if _, ok := needsMap[deleted]; ok {
				if windowMap[deleted] == needsMap[deleted] {
					validsInWindow--
				}
				windowMap[deleted]--
			}
		}
	}

	return answer
}
