package easy557

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	var reverseWord func(string) string
	reverseWord = func(s string) (ans string) {
		for i := len(s) - 1; i >= 0; i-- {
			ans += string(s[i])
		}
		return
	}

	var ans string
	for i, word := range words {
		ans += reverseWord(word)
		if i < len(words)-1 {
			ans += " "
		}
	}

	return ans
}

func reverseWords1(s string) string {
	words := strings.Split(s, " ")

	var reverseWord func(string) string
	reverseWord = func(s string) (ans string) {
		for i := len(s) - 1; i >= 0; i-- {
			ans += string(s[i])
		}
		return
	}

	var ans []string
	for _, word := range words {
		ans = append(ans, reverseWord(word))
	}

	// strings.Join has better performance than string addition operator.
	return strings.Join(ans, " ")
}
