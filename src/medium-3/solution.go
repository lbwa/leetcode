package medium3

// LengthOfLongestSubstring is used to find longest substring without repeating characters
func LengthOfLongestSubstring(s string) int {
	left, right, answer := 0, 0, 0
	windowMap := map[byte]int{}

	for right < len(s) {
		current := s[right]
		right++
		windowMap[current]++

		// 当当前字母频率大于 1 表示存在重复字符，那么缩小窗口
		for windowMap[current] > 1 {
			deleted := s[left]
			left++
			windowMap[deleted]--
		}
		answer = max(answer, right-left)
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
