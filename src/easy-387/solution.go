package easy367

func firstUniqChar(s string) int {
	chars := make(map[rune]int)
	for _, ch := range s {
		chars[ch]++
	}

	for i, ch := range s {
		if chars[ch] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqCharFromSlice(s string) int {
	chars := make([]int, 26)
	for _, ch := range s {
		chars[ch-'a']++
	}

	for i, ch := range s {
		if chars[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}
