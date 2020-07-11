package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	left, right, answer := 0, -1, 0
	uniqueChars := make(map[byte]int)
	length := len(s)

	for left < length {
		// 缩小窗口
		if left != 0 {
			delete(uniqueChars, s[left-1])
		}

		// 增大窗口
		for right+1 < length && !isInMap(uniqueChars, s[right+1]) {
			uniqueChars[s[right+1]] = right + 1
			right++
		}

		answer = max(answer, right-left+1)
		left++
	}

	return answer
}

func isInMap(mapping map[byte]int, key byte) bool {
	_, exist := mapping[key]
	return exist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
