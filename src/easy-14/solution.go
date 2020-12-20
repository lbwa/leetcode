package easy14

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ``
	}

	for i := 0; i < len(strs); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
