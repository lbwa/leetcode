package medium567

// CheckInclusion is used to test permutation in string
func CheckInclusion(s1, s2 string) bool {
	left, right := 0, 0
	needsMap, windowMap := map[byte]int{}, map[byte]int{}
	validsInWindow := 0

	for i := range s1 {
		needsMap[s1[i]]++
	}

	for right < len(s2) {
		current := s2[right]
		right++
		if _, ok := needsMap[current]; ok {
			windowMap[current]++
			if windowMap[current] == needsMap[current] {
				validsInWindow++
			}
		}

		// window 长度不小于 s1 的长度时，表示 s1 可能是 window 区域的异
		// 位词，即 s1 可能是 s2 的子串排列，故缩小窗口
		for (right - left) >= len(s1) {
			if validsInWindow == len(needsMap) {
				return true
			}
			deleted := s2[left]
			left++
			if _, ok := needsMap[deleted]; ok {
				if windowMap[deleted] == needsMap[deleted] {
					validsInWindow--
				}
				windowMap[deleted]--
			}
		}
	}

	return false
}
