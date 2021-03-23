package medium17

import "strings"

func letterCombinations(digits string) []string {
	ans := []string{}

	if len(digits) < 1 {
		return ans
	}

	pendings := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	track := []string{}
	var backtracking func(int)
	backtracking = func(current int) {
		if len(track) == len(digits) {
			ans = append(ans, strings.Join(track, ""))
			return
		}
		for _, char := range pendings[digits[current]] {
			track = append(track, string(char))
			backtracking(current + 1)
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	return ans
}
