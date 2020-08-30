package easy9

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	chars := strconv.Itoa(x)
	length := len(chars)
	for i := 0; i < length; i++ {
		if chars[i] != chars[length-1-i] {
			return false
		}
	}
	return true
}
