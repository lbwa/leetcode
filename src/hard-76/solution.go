package hard76

import "math"

// MinWindow is used to find minimum window substring
func MinWindow(s, t string) string {
	left, right := 0, 0 // 左闭右开
	// needsMap 为目标字符串的字母频率映射
	// windowMap 为窗口字符串的字母频率映射
	needsMap, windowMap := map[byte]int{}, map[byte]int{}
	// 表示满足 needsMap 的条件的字符个数，如果 validsInWindow 与 needsMap 的大小相同
	// 那么说明窗口已经满足条件，即已经完全覆盖 T
	validsInWindow := 0
	answerStart, answerLength := 0, math.MaxInt32

	for i := range t {
		needsMap[t[i]]++
	}

	for right < len(s) {
		// 即将进入窗口的字符
		current := s[right]
		right++ // 增大窗口
		// 如果是目标单字符，那么更新映射中的单个字符频率
		if _, ok := needsMap[current]; ok {
			windowMap[current]++
			// 仅在相同频率时，才会更新 validsInWindow
			if windowMap[current] == needsMap[current] {
				validsInWindow++
			}
		}

		// 即已经存在可覆盖 T 的窗口，此时窗口中对应到 t 中单个字符的频率一定等于 t 中
		// 对应的单个字符的频率
		for validsInWindow == len(needsMap) {
			// 当当前窗口长度小于原有值时，更新最小长度
			if right-left < answerLength {
				answerStart = left
				answerLength = right - left
			}
			// 缩小窗口，为了尽可能找到满足覆盖 t 的最短子串
			deleted := s[left]
			left++
			// 因为缩小了窗口，故要移除对应的字符频率
			if _, ok := needsMap[deleted]; ok {
				if windowMap[deleted] == needsMap[deleted] {
					validsInWindow--
				}
				windowMap[deleted]--
			}
		}
	}

	if answerLength == math.MaxInt32 {
		return ``
	}
	return s[answerStart:(answerStart + answerLength)]
}
