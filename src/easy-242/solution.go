package easy242

// IsAnagram is used to judge anagram
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCodeMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		charCodeMap[s[i]]++
		charCodeMap[t[i]]--
	}
	for charCode := range charCodeMap {
		if charCodeMap[charCode] != 0 {
			return false
		}
	}
	return true
}
